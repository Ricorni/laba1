// Очередь.
// добавление и удаление элемента (push и pop),
// чтение.

#pragma once

#include "node.h"

#include <stdexcept>
#include <iostream>

template <typename T>
struct Queue {
    Node<T>* head = nullptr;
    Node<T>* tail = nullptr;

    Queue() {
        head = nullptr;
        tail = nullptr;
    }

    ~Queue() {
        while (head != nullptr) {
            Node<T>* next = head->next;
            delete head;
            head = next;
        }
        tail = nullptr;
    }

    void push(const T& data) {
        Node<T>* newNode = new Node<T>(data, nullptr);

        if (tail != nullptr) {
            tail->next = newNode;
        } else {
            head = newNode;
        }
        tail = newNode;
    }

    T pop() {
        if (head == nullptr) {
            throw std::out_of_range("queue is empty");
        }

        Node<T>* delNode = head;
        head = delNode->next;
        T data = delNode->data;
        delete delNode;

        if (head == nullptr) {
            tail = nullptr;
        }
        return data;
    }

    T& front() {
        if (head == nullptr) {
            throw std::out_of_range("queue is empty");
        }
        return head->data;
    }

    T& back() {
        if (tail == nullptr) {
            throw std::out_of_range("queue is empty");
        }
        return tail->data;
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