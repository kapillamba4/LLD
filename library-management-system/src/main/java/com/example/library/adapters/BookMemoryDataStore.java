package com.example.library.adapters;

import java.util.Objects;

import com.example.library.models.Book;
import com.example.library.persistence.InMemoryDatabase;
import com.example.library.repository.BookRepository;

public class BookMemoryDataStore extends InMemoryDatabase<Book> implements BookRepository {

    @Override
    public void addBooksToInventory(Book book) {
        Book existingBook = findByKey(book.getIfsn());
        if (Objects.nonNull(existingBook)) {
            Book updatedBook = Book.builder()
                .bookName(existingBook.getBookName())
                .ifsn(existingBook.getIfsn())
                .copies(existingBook.getCopies() + book.getCopies())
                .build();

            save(book.getIfsn(), updatedBook);
            return;
        }
        save(book.getIfsn(), book);
    }

    @Override
    public Book getBookByIfsn(String ifsn) {
        return findByKey(ifsn);
    }

    @Override
    public boolean borrowBook(String ifsn) {
        Book book = findByKey(ifsn);
        if (book == null || book.getCopies() == 0) {
            return false;
        }
        book.setCopies(book.getCopies() - 1);
        save(ifsn, book);
        return true;
    }

    @Override
    public boolean returnBook(String ifsn) {
        Book book = findByKey(ifsn);
        if (book == null) {
            return false;
        }
        book.setCopies(book.getCopies() + 1);
        save(ifsn, book);
        return true;
    }
    
}
