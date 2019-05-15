#!/bin/sh
#
# Communicate with the API using cURL. The draft is sent in the request.

# GEEKMAIL_APITOKEN= # Your API Token from https://geekmail.app/setup
GEEKMAIL_APITOKEN=${GEEKMAIL_APITOKEN:?You must provide your API Token}

TEMPLATE=$(cat <<'END_TEMPLATE' | sed ':a;N;$!ba;s/\n/\\n/g'
To: {{ .To }}
Subject: Hello world!

{{ if .Name }}Hey {{ .Name }},{{ else }}Hey there,{{ end }}

I'm happy that you joined us.

-- 
Author
END_TEMPLATE
)

curl -X POST https://geekmail.app/api/1.0/draft/create \
	-H "Content-Type: application/json" \
	-H "Authorization: Bearer $GEEKMAIL_APITOKEN" \
	-d '{
		"template":'"\"$TEMPLATE\""',
		"vars":{
			"To":"John Doe <john@example.com>",
			"Name":"John"
		},
		"labels":["GeekMail"]
	}'
echo # curl doesn't produce a newline
