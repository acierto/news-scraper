var app = angular.module('NewsScraperApp', ['ngRoute', 'ngResource', 'ngSanitize']);


app.config(function($locationProvider) {
    $locationProvider.html5Mode(false).hashPrefix('!');
});