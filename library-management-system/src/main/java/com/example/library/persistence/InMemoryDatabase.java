package com.example.library.persistence;

import java.util.*;

import com.example.library.config.Constants;

public abstract class InMemoryDatabase<T> {
    protected HashMap<String, Map<String, T>> storage = new HashMap<>();

    private Map<String, T> getStorage(String type) {
        if (!storage.containsKey(type)) {
            storage.put(type, new HashMap<>());
        }
        return storage.get(type);
    }

    public void save(String key, T value) {
        getStorage(Constants.defaultStorageType).put(key, value);
    }

    public T findByKey(String key) {
        return getStorage(Constants.defaultStorageType).get(key);
    }

    public void deleteByKey(String key) {
        getStorage(Constants.defaultStorageType).remove(key);
    }

    public Map<String, T> findAll() {
        return getStorage(Constants.defaultStorageType);
    }

    public void save(String storageType, String key, T value) {
        getStorage(storageType).put(key, value);
    }

    public T findByKey(String storageType, String key) {
        return getStorage(storageType).get(key);
    }

    public void deleteByKey(String storageType, String key) {
        getStorage(storageType).remove(key);
    }

    public Map<String, T> findAll(String storageType) {
        return getStorage(storageType);
    }
}
