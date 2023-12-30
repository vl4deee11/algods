#include <iostream>

struct TreapNode {
    TreapNode* l;
    TreapNode* r;
    int prior;
    int key;

    TreapNode(int k) : l(nullptr), r(nullptr), prior(rand()), key(k) {}
};

TreapNode* searchL(TreapNode* root, int x) {
    if (root == nullptr) return new TreapNode(0);
    if (root->key >= x) return searchL(root->l, x);
    return root;
}

TreapNode* searchR(TreapNode* root, int x) {
    if (root == nullptr) return new TreapNode(0);
    if (root->key <= x) return searchR(root->r, x);
    return root;
}

void split(TreapNode* root, TreapNode** l, TreapNode** r, int x) {
    if (root == nullptr) {
        *l = nullptr;
        *r = nullptr;
        return;
    }
    if (root->key > x) {
        split(root->l, l, &root->l, x);
        *r = root;
    } else {
        split(root->r, &root->r, r, x);
        *l = root;
    }
}

void insert(TreapNode** root, TreapNode* it) {
    if (*root == nullptr) {
        *root = it;
        return;
    }
    if (it->prior >= (*root)->prior) {
        split(*root, &it->l, &it->r, it->key);
        *root = it;
    } else {
        if ((*root)->key <= it->key) {
            insert(&(*root)->r, it);
        } else {
            insert(&(*root)->l, it);
        }
    }
}