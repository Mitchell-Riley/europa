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

#include <stdio.h>
#include <stdlib.h>
#include <dispatch/dispatch.h>
#include <europa/list.h>

list_t* list_new(list_node_t* node, dispatch_queue_t q)
{
	list_t* list = malloc(sizeof(*list));

	if(list)
	{
		list->count = 1;
		list->head = node;
		list->tailp = &list->head;
		if(q)
			dispatch_retain(q);
		else
		{
			/* 36 == string length + 64-bit pointer address + '0' */
			char str[36] = {0};
			snprintf(str, sizeof(str) - 1, "europa.list.new.%p", list);
			q = dispatch_queue_create(str, NULL);
		}
		list->queue = q;
	}

	return list;
}

static void list_free(list_t* list)
{
	list_foreach(list, ^(list_node_t* elem){ free(elem); });
	dispatch_release(list->queue);
	free(list);
	list = NULL;
}

list_t* list_retain(list_t* list)
{
	dispatch_queue_t queue = dispatch_queue_create("europa.list.retain", NULL);
	dispatch_sync(queue, ^{ list->count++; });
	dispatch_release(queue);
}

void list_release(list_t* list)
{
	dispatch_queue_t queue = dispatch_queue_create("europa.list.release", NULL);
	dispatch_sync(queue, ^{
		if(list->count--)
		{
			dispatch_queue_t low = dispatch_get_global_queue(DISPATCH_QUEUE_PRIORITY_LOW, 0);
			dispatch_async(low, ^{ list_free(list); });
		}
	});
	dispatch_release(queue);
}

void list_prepend_node(list_t* list, list_node_t* node)
{
	node->next = list->head;
	dispatch_sync(list->queue, ^{
		list->head = node;
		if(list->tailp == &list->head)
			list->tailp = &node->next;
	});
}

void list_prepend(list_t* list, void* item)
{
	list_node_t* node = alloca(sizeof(*node));
	node->data = item;
	list_prepend_node(list, node);
}

void list_append_node(list_t* list, list_node_t* node)
{
	dispatch_sync(list->queue, ^{
		*list->tailp = node;
		list->tailp = &node->next;
	});
}

void list_append(list_t* list, void* item)
{
	list_node_t* node = malloc(sizeof(*node));
	node->data = item;
	node->next = list->head;
	list_append_node(list, node);
}

void list_remove(list_t* list, list_node_t* node)
{
	list_node_t* current;
	list_node_t** pnp = &list->head;

	while((current = *pnp) != NULL)
	{
		if(current == node)
		{
			*pnp = node->next;
			dispatch_sync(list->queue, ^{
				if(list->tailp == &node->next)
					list->tailp = pnp;
				node->next = NULL;
			});
			break;
		}
		pnp = &current->next;
	}
}

void list_foreach(list_t* list, void (^yield)(list_node_t*))
{
	dispatch_sync(list->queue, ^{
		list_node_t* current = list->head;

		while(current != NULL)
		{
			yield(current);
			current = current->next;
		}
	});
}
