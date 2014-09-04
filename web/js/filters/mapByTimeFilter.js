angular.module('NewsScraperApp')
    .filter('mapByTime', function () {
        return _.memoize(function (articles) {
            var mappedArticles = [];
            _.each(articles, function(article) {
                mappedArticles.push({"key": new Date(article.Time), "value": article})
            });

            return mappedArticles;
        });
    });