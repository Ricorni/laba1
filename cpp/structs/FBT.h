#pragma once

#include <iostream>
#include <stdexcept>
#include <vector>
#include <string>
#include "node.h"

using namespace std;

template <typename T>
struct FullBinaryTree {
    TreeNode<T>* root = nullptr;

    FullBinaryTree() { root = nullptr; }

    ~FullBinaryTree() {
        clear(root);
    }

    void clear(TreeNode<T>* node) {
        if (node == nullptr) return;
        clear(node->left);
        clear(node->right);
        delete node;
    }

    void insert(const T& data) {
        root = insertNode(root, data);
    }

    TreeNode<T>* insertNode(TreeNode<T>* node, const T& data) {
        if (node == nullptr) return new TreeNode<T>(data);

        if (data < node->data) {
            node->left = insertNode(node->left, data);
        } else if (data > node->data) {
            node->right = insertNode(node->right, data);
        }
        return node;
    }

    bool find(const T& data) const {
        TreeNode<T>* current = root;
        while (current != nullptr) {
            if (data == current->data) return true;
            else if (data < current->data) current = current->left;
            else current = current->right;
        }
        return false;
    }

    bool isFull() const {
        return checkFull(root);
    }

    bool checkFull(TreeNode<T>* node) const {
        if (node == nullptr) return true;
        if (node->left == nullptr && node->right == nullptr) return true;
        if (node->left != nullptr && node->right != nullptr) {
            return checkFull(node->left) && checkFull(node->right);
        }
        return false;
    }

    void print() const {
        if (root == nullptr) {
            cout << "(empty)" << endl;
            return;
        }
        printNode(root, "", false);
    }

    void printNode(TreeNode<T>* node, string prefix, bool isLeft) const {
        if (node == nullptr) return;

        cout << prefix;
        cout << (isLeft ? "|-- " : "\\-- ");
        cout << node->data << endl;

        string newPrefix = prefix + (isLeft ? "|   " : "    ");

        if (node->left || node->right) {
            if (node->left) printNode(node->left, newPrefix, true);
            else cout << newPrefix << "|-- [null]" << endl;

            if (node->right) printNode(node->right, newPrefix, false);
            else cout << newPrefix << "\\-- [null]" << endl;
        }
    }

    // Обход в ширину (BFS) — по уровням. 
    void printBFS() const {
        if (root == nullptr) return;
        
        std::vector<TreeNode<T>*> q;
        q.push_back(root);
        size_t head = 0;

        while (head < q.size()) {
            TreeNode<T>* cur = q[head++];
            cout << cur->data << " ";
            if (cur->left) q.push_back(cur->left);
            if (cur->right) q.push_back(cur->right);
        }
        cout << endl;
    }

    void save(ostream& os) const {
        savePreOrder(root, os);
    }

    void savePreOrder(TreeNode<T>* node, ostream& os) const {
        if (node == nullptr) return;
        os << node->data << endl;
        savePreOrder(node->left, os);
        savePreOrder(node->right, os);
    }
};