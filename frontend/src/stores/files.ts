import { writable } from 'svelte/store';
import type { FileRec } from '$lib/utils/types';

export const files = writable<FileRec[]>([]);
