package com.example.library.usecase;

import com.example.library.dependency.DI;
import com.example.library.models.Book;

public class BooksManager {
    private DI diContainer;

    public BooksManager(DI di) {
        this.diContainer = di;
    }

    // Adds new book entry or update copies of already registered book
    public void addNewBook(String ifsn, String bookName, int copies) {
        diContainer.getBookRepository()
                .addBooksToInventory(Book.builder().bookName(bookName).copies(copies).ifsn(ifsn).build());
    }
}
