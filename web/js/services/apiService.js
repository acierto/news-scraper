angular.module('NewsScraperApp')
    .factory('ApiService', function ($http) {
        return {
            getArticles: function () {
                return $http.get('/read-articles', {cache: true});
            }
        }
    });