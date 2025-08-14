FROM nginx:alpine
COPY default.conf /etc/nginx/conf.d/default.conf
COPY --from=frontend /app/build /usr/share/nginx/html
