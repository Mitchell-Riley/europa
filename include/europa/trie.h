/*******************************************************************
 * Europa Programming Language 
 * Copyright (C) 2010, Jeremy Tregunna, All Rights Reserved.
 *
 * This software project, which includes this module, is protected
 * under Canadian copyright legislation, as well as international
 * treaty. It may be used only as directed by the copyright holder.
 * The copyright holder hereby consents to usage and distribution
 * based on the terms and conditions of the MIT license, which may
 * be found in the LICENSE.MIT file included in this distribution.
 *******************************************************************
 * Project: Europa Programming Language 
 * File: trie.h
 * Description: Definitions for a PATRICIA trie.
 ******************************************************************/

#ifndef __EUROPA__TRIE_H__
#define __EUROPA__TRIE_H__

#include <sys/types.h>
#include <dispatch/dispatch.h>

/*
 * This module defines a PATRICIA trie, which we use for storing
 * our slot table on actors.
 */
 
/* Node structure */
typedef struct trie_node_s
{
	int bit_index;
	char* key;
	void* data;
	struct trie_node_s* left;
	struct trie_node_s* right;
} trie_node_t;

/* Trie itself */

typedef struct trie_s
{
	trie_node_t* head;
	dispatch_queue_t queue;
	unsigned int count;
} trie_t;

#define trie_node_get_left(node) ((node)->left)
#define trie_node_get_right(node) ((node)->right)
#define trie_node_get_key(node) ((node)->key)
#define trie_node_get_data(node) ((node)->data)

extern trie_node_t* trie_node_new(const char*, void*, int, trie_node_t*, trie_node_t*);
extern void trie_node_free(trie_node_t*);
extern void trie_node_set_data(trie_node_t*, void*, size_t);

/* Creates a new trie. The queue will be used to synchronize
   various operations. If the supplied queue is NULL, a new one
   will be created. */
extern trie_t* trie_new(dispatch_queue_t);

/* Increases the reference count by one. */
extern trie_t* trie_retain(trie_t*);

/* Decrements the reference count by one. If the count is 0,
   fires off an async message to free the trie. */
extern void trie_release(trie_t*);

/* Insert a key->value pair into the trie. */
extern trie_node_t* trie_insert(trie_t*, const char*, void*);

/* Find an entry in the trie by the supplied name. */
extern void* trie_lookup(trie_t*, const char*);

/* Find the node with the supplied name. */
extern trie_node_t* trie_lookup_node(trie_t*, const char*);

/* Remove an item from the trie. */
extern int trie_delete(trie_t*, const char*);

#endif /* !__EUROPA__TRIE_H__ */
