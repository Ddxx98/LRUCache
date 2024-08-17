package controllers

import (
	"container/list"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type CacheItem struct {
	Key        string
	Value      string
	Expiration time.Time
}

type LRUCache struct {
	capacity int
	mutex    sync.Mutex
	cache    map[string]*list.Element
	lruList  *list.List
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element, capacity),
		lruList:  list.New(),
	}
}

var cache = NewLRUCache(1024)

func (c *LRUCache) Set(key string, value string, expiration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.cache[key]; ok {
		item := elem.Value.(*CacheItem)
		item.Value = value
		item.Expiration = time.Now().Add(expiration)

		c.lruList.MoveToFront(elem)
	} else {
		if len(c.cache) >= c.capacity {
			c.removeOldest()
		}

		item := &CacheItem{
			Key:        key,
			Value:      value,
			Expiration: time.Now().Add(expiration),
		}
		elem := c.lruList.PushFront(item)
		c.cache[key] = elem
	}
}

func (c *LRUCache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.cache[key]; ok {
		item := elem.Value.(*CacheItem)
		if item.Expiration.After(time.Now()) {
			c.lruList.MoveToFront(elem)
			return item.Value, true
		} else {
			c.removeElement(elem)
		}
	}
	return "nil", false
}

func (c *LRUCache) removeOldest() {
	elem := c.lruList.Back()
	if elem != nil {
		c.removeElement(elem)
	}
}

func (c *LRUCache) removeElement(e *list.Element) {
	item := e.Value.(*CacheItem)
	delete(c.cache, item.Key)
	c.lruList.Remove(e)
}

func (c *LRUCache) Delete(key string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if elem, ok := c.cache[key]; ok {
		c.removeElement(elem)
		return nil
	}
	return fmt.Errorf("key %s not found", key)
}

func SetHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Key               string `json:"key"`
		Value             string `json:"value"`
		ExpirationSeconds int    `json:"expirationSeconds"`
	}
	decoder := json.NewDecoder(r.Body)
	fmt.Println(r.Body, decoder)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	cache.Set(data.Key, data.Value, time.Duration(data.ExpirationSeconds)*time.Second)
	w.WriteHeader(http.StatusOK)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, found := cache.Get(key)
	if !found {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	response := struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}{key, value}
	fmt.Println(response)
	json.NewEncoder(w).Encode(response)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	val := cache.Delete(key)
	if val != nil {
		json.NewEncoder(w).Encode(val)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func Visible(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Connected")
}
