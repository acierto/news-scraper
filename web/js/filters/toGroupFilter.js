angular.module('NewsScraperApp')
    .filter('toGroup', function () {
        return _.memoize(function (articles) {
            var groupedArticles = [];

            var subArray = [];
            _.each(articles, function(article, index) {
                if (index % 3 == 0 && index != 0) {
                    groupedArticles.push(subArray);
                    subArray = [];
                }
                subArray.push(article);
            });

            if (!_.isEmpty(subArray)) {
                groupedArticles.push(subArray);
            }

            return groupedArticles;
        });
    });