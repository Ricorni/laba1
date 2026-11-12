// Стек.
// добавление и удаление элемента (push и pop),
// чтение.

#pragma once

#include "node.h"

#include <stdexcept>
#include <iostream>

template <typename T>
struct Stack {
    Node<T>* head = nullptr;

    Stack() {
        head = nullptr;
    }

    ~Stack() {
        while (head != nullptr) {
            Node<T>* next = head->next;
            delete head;
            head = next;
        }
    }

    void push(const T& data) {
        Node<T>* newNode = new Node<T>(data, head);
        head = newNode;
    }

    T pop() {
        if (head == nullptr) {
            throw std::out_of_range("stack is empty");
        }

        Node<T>* delNode = head;
        head = delNode->next;
        T data = delNode->data;
        delete delNode;
        return data;
    }

    T& front() {
        if (head == nullptr) {
            throw std::out_of_range("stack is empty");
        }
        return head->data;
    }

    bool isEmpty() const {
        return head == nullptr;
    }

    void print() const {
        Node<T>* cur = head;
        while (cur != nullptr) {
            std::cout << cur->data << " ";
            cur = cur->next;
        }
        std::cout << "\n";
    }

    void save(std::ostream& os) const {
        Node<T>* cur = head;
        while (cur != nullptr) {
            os << cur->data << "\n";
            cur = cur->next;
        }
    }
};