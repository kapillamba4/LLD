package com.example.library.adapters;

import com.example.library.config.Constants;
import com.example.library.models.User;
import com.example.library.persistence.InMemoryDatabase;
import com.example.library.repository.UserRepository;

public class UserMemoryDataStore extends InMemoryDatabase<User> implements UserRepository {

    @Override
    public void registerNewMember(User user) {
        save(user.getMemberId(), user);
        save(Constants.contactIndexUserStore, user.getContactNumber(), user);
    }

    @Override
    public User getMemberDetailsByID(String memberId) {
        return findByKey(memberId);
    }

    @Override
    public User getMemberDetailsByContact(String contact) {
        return findByKey(Constants.contactIndexUserStore, contact);
    }
}
