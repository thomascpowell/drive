import { writable } from 'svelte/store';
import type { Res } from '$lib/types';

export const status = writable({} as Res);
