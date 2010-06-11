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

#define EUROPA_OBJ_HEADER		\
	dispatch_queue_t queue;		\
	trie_t* blackboards;		\
	actor_react_f react;		\
	list_t* protos;			\
	char* types;			\
	trie_t* slots;			\
	unsigned int done_lookup:1;	\
	unsigned int activatable:1;	\
	unsigned int marked:1;

typedef struct obj_s
{
	dispatch_queue_t queue;
	trie_t* blackboards;
	actor_react_f react;

	EUROPA_OBJ_HEADER;
} obj_t;

#endif /* !__EUROPA__OBJECT_H__ */
