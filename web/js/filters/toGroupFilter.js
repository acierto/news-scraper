angular.module('NewsScraperApp')
    .filter('toGroup', function () {
        return _.memoize(function (articles) {
            var collectedArticles = [];

            var subArray = [];
            _.each(articles, function(article, index) {
                if (index % 3 == 0 && index != 0) {
                    collectedArticles.push(subArray);
                    subArray = [];
                }
                subArray.push(article);
            });

            if (!_.isEmpty(subArray)) {
                collectedArticles.push(subArray);
            }

            return collectedArticles;
        });
    });