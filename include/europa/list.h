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
 * File: list.h
 * Description: Defines a linked list.
 ******************************************************************/

#ifndef __EUROPA__LIST_H__
#define __EUROPA__LIST_H__

#include <dispatch/dispatch.h>
 
typedef struct list_node_s
{
	struct list_node_s* next;
	void* data;
} list_node_t;

/* The base list type. */
typedef struct list_s
{
	list_node_t* head;
	list_node_t** tailp;
	dispatch_queue_t queue;
	unsigned int count;
} list_t;

/* Creates a new list, setting the head to the supplied node.
   Node can be NULL. Also sets up the dispatch queue, which can also
   be NULL; if NULL, we will create a new serial queue. This queue
   should be a serial queue, and yours should be too. If it is not,
   then this is not thread safe. */
extern list_t* list_new(list_node_t*, dispatch_queue_t);

/* Increase the retain count by one. */
extern list_t* list_retain(list_t*);

/* Decrement the retain count by one, and if zero, asynchronously
   release the resources. */
extern void list_release(list_t*);

/* Add an item to the beginning of list. */
extern void list_prepend(list_t*, void*);
extern void list_prepend_node(list_t*, list_node_t*);

/* Add an item to the end of a list. */
extern void list_append(list_t*, void*);
extern void list_append_node(list_t*, list_node_t*);

/* Remove a node from the list. */
extern void list_remove(list_t*, list_node_t*);

/* Iterate the list. Executes the supplied block for each node. */
extern void list_foreach(list_t*, void (^)(list_node_t*));

#endif /* !__EUROPA__LIST_H__ */
