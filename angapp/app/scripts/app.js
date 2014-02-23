'use strict';

angular.module('angappApp', [
  'ngCookies',
  'ngResource',
  'ngSanitize',
  'ngRoute'
])
  .config(function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'views/main.html',
        controller: 'MainCtrl'
      })
      .when('/about', {
        templateUrl: 'views/about.html',
        controller: 'AboutCtrl'
      })
      .when('/persons', {
        templateUrl: 'views/persons.html',
        controller: 'PersonsCtrl'
      })
      .when('/persons/new', {
        templateUrl: 'views/persons/new.html',
        controller: 'NewPersonsCtrl'
      })
      .when('/persons/:id', {
        templateUrl: 'views/persons/show.html',
        controller: 'PersonsShowCtrl'
      })
      .when('/persons/edit/:id', {
        templateUrl: 'views/persons/edit.html',
        controller: 'EditPersonsCtrl'
      })
      .otherwise({
        redirectTo: '/'
      });
  });
