import { writable } from 'svelte/store';
import type { Res } from '$lib/utils/types';

export const status = writable({} as Res);
