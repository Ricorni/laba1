#pragma once

template <typename T>
struct Node {
    T data;
    Node* next;

    Node() : data(T()), next(nullptr) {} 

    Node(const T& data) : data(data), next(nullptr) {}

    Node(const T& data, Node* next) : data(data), next(next) {}
};

template <typename T>
struct DoubleNode {
    T data;
    DoubleNode* next;
    DoubleNode* prev;

    DoubleNode() : data(T()), next(nullptr), prev(nullptr) {} 

    DoubleNode(const T& data) : data(data), next(nullptr), prev(nullptr) {}

    DoubleNode(const T& data, DoubleNode* next, DoubleNode* prev) 
        : data(data), next(next), prev(prev) {}
};

template <typename T>
struct TreeNode {
    T data;
    TreeNode<T>* left;
    TreeNode<T>* right;

    TreeNode() : data(T()), left(nullptr), right(nullptr) {} 
    TreeNode(const T& val) : data(val), left(nullptr), right(nullptr) {}
}; 