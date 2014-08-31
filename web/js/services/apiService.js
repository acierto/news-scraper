angular.module('NewsScraperApp')
    .factory('ApiService', function ($http) {
        return {
            getArticles: function () {
                return $http.get('/news-scraper/web/json/articles.json', {cache: true});
            }
        }
    });