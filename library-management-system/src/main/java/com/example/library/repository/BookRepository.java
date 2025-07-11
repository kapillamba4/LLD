package com.example.library.repository;

import com.example.library.models.Book;

public interface BookRepository {
    // Add book along with copies to the inventory along with existing books on shelf
    void addBooksToInventory(Book book);
    // Get book details by IFSN
    Book getBookByIfsn(String ifsn);
    // Borrow book
    boolean borrowBook(String ifsn);
    // Return book
    boolean returnBook(String ifsn);
}
