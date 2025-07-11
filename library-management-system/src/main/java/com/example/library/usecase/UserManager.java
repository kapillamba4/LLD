package com.example.library.usecase;

import java.util.UUID;

import com.example.library.dependency.DI;
import com.example.library.models.User;

public class UserManager {
    private DI diContainer;

    public UserManager(DI di) {
        this.diContainer = di;
    }

    public void addNewMember(String name, String contact, String address) {
        User user = User.builder().name(name).contactNumber(contact).address(address)
                .memberId(UUID.randomUUID().toString()).build();
        diContainer.getUserRepository().registerNewMember(user);
    }

    public User fetchMemberDetailsByID(String memberId) {
        return diContainer.getUserRepository().getMemberDetailsByID(memberId);
    }

    public User fetchMemberDetailsByContact(String contact) {
        return diContainer.getUserRepository().getMemberDetailsByContact(contact);
    }

}
