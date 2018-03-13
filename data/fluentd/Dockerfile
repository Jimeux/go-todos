FROM fluent/fluentd:v0.12-onbuild

RUN apk add --update --virtual .build-deps \
      sudo build-base zlib-dev postgresql-dev ruby-dev \
  && sudo gem install fluent-plugin-redshift \
  && sudo gem sources --clear-all \
  && apk del .build-deps
RUN rm -rf /var/cache/apk/* \
      /home/fluent/.gem/ruby/2.3.0/cache/*.gem
