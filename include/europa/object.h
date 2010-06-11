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
 * File: 
 * Description: 
 ******************************************************************/

#ifndef __EUROPA__OBJECT_H__
#define __EUROPA__OBJECT_H__
 
#include <dispatch/dispatch.h>
#include <europa/list.h>
#include <europa/trie.h>

/**
 * Objects are modelled on the concept of actors, and are the same
 * structure for all objects. As a result, the only method of
 * communicating with an object is through messages. However, actors
 * create blackboards, and ask their intended receiver to subscribe
 * to it, and fire the message off.
 */
typedef struct obj_s
{
	/* Messages are queued up in this queue and acted on according
	   to the custom "react" function, or the default. */
	dispatch_queue_t queue;

	/* Stores key->value pairs for any blackboards we push
	   messages onto. */
	trie_t* blackboards;

	/* List of obj_t* types which act as parents in the inheritance
	   tree. */
	list_t* protos;

	/* Reserved for the future. */
	char* types;

	/* Key->value pairs representing our slot table. */
	trie_t* slots;

	/* Set if we have already looked at this object when working
	   through the inheritance graph. */
	unsigned int done_lookup:1;

	/* Set if the object will implicitly call its activate function
	   when looked up. */
	unsigned int activatable:1;
} obj_t;

#endif /* !__EUROPA__OBJECT_H__ */
