import { writable } from 'svelte/store';
import type { File } from '$lib/utils/types';

export const files = writable<File[]>([]);
