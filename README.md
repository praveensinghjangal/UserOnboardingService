Module -> Onboarding_Service
Application Run Command -> go run main.go
Below are the Curls for Apis :-
a.) User signUp -> curl --location 'localhost:8080/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"jangalpraveen@gmail.com",
    "password":"1234567Ja$$$"
}'

b.) User signin -> curl --location 'localhost:8080/signin' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"jangalpraveen@gmail.com",
    "password":"1234567Ja$$$"
}'

c.) User Revoke or logout from the Application -> curl --location --request POST 'localhost:8080/api/logout' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiIsImV4cCI6MTczMjc5MjIwNX0.TmdVnnI5bqFuDcJrn55iJaZ07E5B5ms8-CgC8T-jpAc'

d.) User wants to Refresh the token before its expiry -> curl --location --request POST 'localhost:8080/api/refresh' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiIsImV4cCI6MTczMjc5MjE3Mn0.pAPeD8HBDFVXhTz6TOeVfEeDZZsT9G7-ah1O5H2sums'
