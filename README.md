# Table of Contents
- [Table of Contents](#table-of-contents)
- [Introduction](#introduction)
- [Full-stack Applications](#full-stack-applications)
  - [E-commerce (shopping cart)](#e-commerce-shopping-cart)
    - [Server side implementations](#server-side-implementations)
    - [Client side implementations](#client-side-implementations)
  - [Blog/CMS](#blogcms)
    - [Server side implementations](#server-side-implementations-1)
    - [Client side](#client-side)
      - [The next come are](#the-next-come-are)
  - [Simple CRUD(Create, Read, Update, Delete)](#simple-crudcreate-read-update-delete)
    - [Server side implementations](#server-side-implementations-2)
    - [Client side implementations](#client-side-implementations-1)
      - [The next come are](#the-next-come-are-1)
  - [CRUD + Pagination](#crud--pagination)
    - [Server side implementations](#server-side-implementations-3)
      - [The next come are](#the-next-come-are-2)
    - [Client side implementations](#client-side-implementations-2)
      - [The next come are](#the-next-come-are-3)
- [Social media links](#social-media-links)
- [Commands used to build the project](#commands-used-to-build-the-project)
- [Follow me](#follow-me)
    
# Introduction
This is a mini application implementing REST CRUD operations using Go with Beego(Web and ORM).

# Getting Started
- Rename .env.example to .env and set the database settings for your Database(MySQL works, PSQL maybe, SQLite does not work)
- Install Beego Go package. Install Beego
`go get -u github.com/astaxie/beego`
- Install Bee (Optional). It is a tool to scaffold a Beego project
`go get -u github.com/beego/bee`
- Generate Swagger
`bee run -downdoc=true -gendoc=true`
To be honest I did not use Bee... It is not needed.

# Full-stack Applications
## Simple Crud
### Server side implementations
- [Python Django + Rest Framework](https://github.com/melardev/DjangoRestFrameworkCrud)
- [Python Django](https://github.com/melardev/DjanogApiCrud)
- [Python Flask](https://github.com/melardev/FlaskApiCrud)
- [Asp.Net Web Api 2](https://github.com/melardev/AspNetWebApiCrud)
- [Asp.Net Core](https://github.com/melardev/AspNetCoreApiCrud)
- [Asp.Net Core + MediatR](https://github.com/melardev/AspNetCoreApiCrudMediatR)
- [Laravel](https://github.com/melardev/LaravelApiCrud)
- [Ruby On Rails](https://github.com/melardev/RailsApiCrud)
- [Ruby On Rails + JBuilder](https://github.com/melardev/RailsApiJBuilderCrud)
- [Spring Boot + Spring Data JPA](https://github.com/melardev/SpringBootApiJpaCrud)
- [Kotlin Spring Boot + Spring Data JPA](https://github.com/melardev/KotlinSpringBootApiJpaCrud)
- [Spring Boot + JAX-RS(Jersey) + Spring Data JPA](https://github.com/melardev/SpringBootApiJerseySpringDataCrud)
- [Spring Boot Reactive + MongoDB Reactive](https://github.com/melardev/SpringBootApiReactiveMongoCrud)
- [Go + GORM](https://github.com/melardev/GoGinGonicApiGormCrud)
- [Go + Gorilla + GORM](https://github.com/melardev/GoMuxGormApiCrud)
- [Go + Beego(Web and ORM)](https://github.com/melardev/GoBeegoApiCrud)
- [Express.JS + Sequelize ORM](https://github.com/melardev/ExpressSequelizeApiCrud)
- [Express.JS + BookShelf ORM](https://github.com/melardev/ExpressBookshelfApiCrud)
- [Express.JS + Mongoose](https://github.com/melardev/ExpressMongooseApiCrud)

#### Microservices
- [Java Spring Boot Zuul + Rest](https://github.com/melardev/JavaSpringBootZuulRestApiCrud)
- [Kotlin Spring Boot Zuul + Rest](https://githubcom/melardev/KotlinSpringBootZuulRestApiCrud)

### Client side implementations
- [React](https://github.com/melardev/ReactCrudAsync)
- [React + Redux](https://github.com/melardev/ReactReduxAsyncCrud)
- [Angular](https://github.com/melardev/AngularApiCrud)
- [Vue](https://github.com/melardev/VueAsyncCrud)
- [Vue + Vuex](https://github.com/melardev/VueVuexAsyncCrud)

#### The next come are
- Angular NgRx-Store
- Angular + Material
- React + Material
- React + Redux + Material
- Vue + Material
- Vue + Vuex + Material
- Ember
- Vanilla javascript


## Crud + Pagination
### Server side implementations
- [Spring Boot + Spring Data + Jersey](https://github.com/melardev/SpringBootJerseyApiPaginatedCrud)
- [Spring Boot + Spring Data](https://github.com/melardev/SpringBootApiJpaPaginatedCrud)
- [Spring Boot Reactive + Spring Data Reactive](https://github.com/melardev/ApiCrudReactiveMongo)
- [Go with Gin Gonic](https://github.com/melardev/GoGinGonicApiPaginatedCrud)
- [Go with Gorilla + GORM](https://github.com/melardev/GoMuxGormApiCrudPagination)
- [Go + Beego(Web and ORM)](https://github.com/melardev/GoBeegoApiCrudPagination)
- [Laravel](https://github.com/melardev/LaravelApiPaginatedCrud)
- [Rails + JBuilder](https://github.com/melardev/RailsJBuilderApiPaginatedCrud)
- [Rails](https://github.com/melardev/RailsApiPaginatedCrud)
- [NodeJs Express + Sequelize](https://github.com/melardev/ExpressSequelizeApiPaginatedCrud)
- [NodeJs Express + Bookshelf](https://github.com/melardev/ExpressBookshelfApiPaginatedCrud)
- [NodeJs Express + Mongoose](https://github.com/melardev/ExpressApiMongoosePaginatedCrud)
- [Python Django](https://github.com/melardev/DjangoApiCrudPaginated)
- [Python Django + Rest Framework](https://github.com/melardev/DjangoRestFrameworkPaginatedCrud)
- [Python Flask](https://github.com/melardev/FlaskApiPaginatedCrud)
- [AspNet Core](https://github.com/melardev/AspNetCoreApiPaginatedCrud)
- [AspNet Web Api 2](https://github.com/melardev/WebApiPaginatedAsyncCrud)

#### MicroServices
- [Java Spring Boot Zuul + Rest](https://github.com/melardev/JavaSpringBootZuulRestApiPaginatedCrud)
- [Kotlin Spring Boot Zuul + Rest](https://github.com/melardev/KotlinSpringBootZuulRestApiPaginatedCrud)

#### The next come are
- NodeJs Express + Knex
- Flask + Flask-Restful
- Laravel + Fractal
- Laravel + ApiResources
- Go with Mux
- AspNet Web Api 2
- Jersey
- Elixir

### Client side implementations
- [Angular](https://github.com/melardev/AngularPaginatedAsyncCrud)
- [React-Redux](https://github.com/melardev/ReactReduxPaginatedAsyncCrud)
- [React](https://github.com/melardev/ReactAsyncPaginatedCrud)
- [Vue + Vuex](https://github.com/melardev/VueVuexPaginatedAsyncCrud)
- [Vue](https://github.com/melardev/VuePaginatedAsyncCrud)


#### The next come are
- Angular NgRx-Store
- Angular + Material
- React + Material
- React + Redux + Material
- Vue + Material
- Vue + Vuex + Material
- Ember
- Vanilla javascript

## E-commerce
### Server side implementations
- [Spring Boot + Spring Data Hibernate](https://github.com/melardev/SBootApiEcomMVCHibernate)
- [Spring Boot + JAX-RS Jersey + Spring Data Hibernate](https://github.com/melardev/SpringBootEcommerceApiJersey)
- [Node Js + Sequelize](https://github.com/melardev/ApiEcomSequelizeExpress)
- [Node Js + Bookshelf](https://github.com/melardev/ApiEcomBookshelfExpress)
- [Node Js + Mongoose](https://github.com/melardev/ApiEcomMongooseExpress)
- [Python Django](https://github.com/melardev/DjangoRestShopApy)
- [Flask](https://github.com/melardev/FlaskApiEcommerce)
- [Golang go gonic](https://github.com/melardev/api_shop_gonic)
- [Ruby on Rails](https://github.com/melardev/RailsApiEcommerce)
- [AspNet Core](https://github.com/melardev/ApiAspCoreEcommerce)
- [Laravel](https://github.com/melardev/ApiEcommerceLaravel)

The next to come are:
- Spring Boot + Spring Data Hibernate + Kotlin
- Spring Boot + Jax-RS Jersey + Hibernate + Kotlin
- Spring Boot + mybatis
- Spring Boot + mybatis + Kotlin
- Asp.Net Web Api v2
- Elixir
- Golang + Beego
- Golang + Iris
- Golang + Echo
- Golang + Mux
- Golang + Revel
- Golang + Kit
- Flask + Flask-Restful
- AspNetCore + NHibernate
- AspNetCore + Dapper

### Client side implementations
This client side E-commerce application is also implemented using other client side technologies:
- [React Redux](https://github.com/melardev/ReactReduxEcommerceRestApi)
- [React](https://github.com/melardev/ReactEcommerceRestApi)
- [Vue](https://github.com/melardev/VueEcommerceRestApi)
- [Vue + Vuex](https://github.com/melardev/VueVuexEcommerceRestApi)
- [Angular](https://github.com/melardev/AngularEcommerceRestApi)

## Blog/CMS
### Server side implementations
### Client side
#### The next come are
- Angular NgRx-Store
- Angular + Material
- React + Material
- React + Redux + Material
- Vue + Material
- Vue + Vuex + Material
- Ember

# Social media links
- [Youtube Channel](https://youtube.com/melardev) I publish videos mainly on programming
- [Blog](http://melardev.com) Sometimes I publish the source code there before Github
- [Twitter](https://twitter.com/@melardev) I share tips on programming
- [Instagram](https://instagram.com/melar_dev) I share from time to time nice banners

# Resources
- [Beego Intro](https://beego.me/docs/intro/)
- [Beego ORM Query](https://beego.me/docs/mvc/model/query.md)
- [Beego Config](https://beego.me/docs/mvc/controller/config.md)