#pragma once

#include "structs/stack.h"
#include "structs/queue.h"
#include "structs/forward_list.h" 
#include "structs/doubly_list.h" 
#include "structs/array_list.h"    
#include "structs/FBT.h"           

#include <map>
#include <string>

enum class Command {
    // Массив
    MPUSH, MDEL, MGET, MINSERT, MSET, MSIZE,

    // Односвязный список
    FINSERTAFTER, FINSERTBEFORE, FREMOVEAFTER, FREMOVEBEFORE,
    FPUSH, FPUSHFRONT, 
    FPOPFRONT, FPOPBACK, FDEL,
    FGET, FGETBACK, FGETIDX, FFIND,

    // Двусвязный список
    LINSERTAFTER, LINSERTBEFORE, LREMOVEAFTER, LREMOVEBEFORE,
    LPUSH, LPUSHFRONT, 
    LPOPFRONT, LPOPBACK, LDEL,
    LGET, LGETBACK, LGETIDX, LFIND,

    // Очередь
    QPUSH, QPOP, QGET, QBACK,

    // Стек
    SPUSH, SPOP, SGET,

    // Дерево
    TINSERT, TDEL, TGET, TISFULL, TBFS,

    // Общее
    PRINT,

    UNKNOWN
};

struct Database {
    std::map<std::string, Stack<std::string>> stacks;
    std::map<std::string, Queue<std::string>> queues;
    
    std::map<std::string, flist<std::string>> forwardLists; 
    std::map<std::string, DoubleList<std::string>> doublyLists; 
    std::map<std::string, ArrayList<std::string>> arrays; 
    std::map<std::string, FullBinaryTree<int>> trees; 
};