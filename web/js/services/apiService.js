angular.module('NewsScraperApp')
    .factory('ApiService', function ($http) {
        return {
            getArticles: function () {
                return $http.get('/json/articles.json', {cache: true});
            }
        }
    });