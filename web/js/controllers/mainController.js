angular.module('NewsScraperApp').controller('MainController',
    function ($scope, ApiService) {

        ApiService.getArticles().then(function(articles) {
            $scope.newsSources = articles.data;
        });

        $scope.hasImage = function(article) {
            return article.Img != '';
        };

        $scope.openLink = function(link) {
            $("#external-article").html('<object data="' + link + '"/>');
        }
    });