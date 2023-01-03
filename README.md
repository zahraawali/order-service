# Order Micro-service

## Introduction

I used both [auth-service](https://github.com/cloud-1401-project/auth-service) and [product-service](https://github.com/cloud-1401-project/product_service) micro-services along with [order micro-service](https://github.com/zahraawali/order-service) that I implemented with [MariaDB](https://hub.docker.com/_/mariadb) database. I made 4 routes as follows:

- GET `/cart` - used for retreive all cart items for the authenticated user.
- POST `/cart` - used for add new item to the cart for the authenticated user.
- DELETE `/cart` - used to empty the cart for the authenticated user.
- DELETE `/cart/{product_id}` - used to delete a specific item in the cart for the authenticated user.

## In-detail Explaination

> Note: In this micro-service I've used `MariaDB` database for adding, modifying and deleting items from cart table.

First of all, I made a `router` variable that is instance of gin `Engine` as I used gin framework. As well as a middleware for verify token along with the request using gRPC, and save `user_id` into gin `request context` therefore I can use it later `(in the handler)`.
After request reaches the handler, that means the user is authenticated and the token is valid.

Also, as I mentioned that I used a `database`, I made use of `gorm` ORM database helper so that I can use it with ORM way.
