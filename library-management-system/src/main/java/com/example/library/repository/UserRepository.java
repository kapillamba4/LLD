package com.example.library.repository;

import com.example.library.models.User;

public interface UserRepository {
    void registerNewMember(User user);
    User getMemberDetailsByID(String memberId);
    User getMemberDetailsByContact(String contact);
}
