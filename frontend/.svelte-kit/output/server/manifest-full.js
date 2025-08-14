export const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["robots.txt"]),
	mimeTypes: {".txt":"text/plain"},
	_: {
		client: {start:"_app/immutable/entry/start.Cbck53Ca.js",app:"_app/immutable/entry/app.Bbtd7YyX.js",imports:["_app/immutable/entry/start.Cbck53Ca.js","_app/immutable/chunks/DxsrY6WV.js","_app/immutable/chunks/DsF5PmS0.js","_app/immutable/chunks/CtJeomz_.js","_app/immutable/chunks/CTjujSVZ.js","_app/immutable/entry/app.Bbtd7YyX.js","_app/immutable/chunks/CtJeomz_.js","_app/immutable/chunks/DsF5PmS0.js","_app/immutable/chunks/CTjujSVZ.js","_app/immutable/chunks/DsnmJJEf.js"],stylesheets:[],fonts:[],uses_env_dynamic_public:false},
		nodes: [
			__memo(() => import('./nodes/0.js')),
			__memo(() => import('./nodes/1.js')),
			__memo(() => import('./nodes/2.js'))
		],
		remotes: {
			
		},
		routes: [
			{
				id: "/",
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 2 },
				endpoint: null
			}
		],
		prerendered_routes: new Set([]),
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})();
