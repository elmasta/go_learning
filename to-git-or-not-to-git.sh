curl -s https://zone01normandie.org/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"ValentinR\"}}){id}}"}' | jq -r .data.user[].id