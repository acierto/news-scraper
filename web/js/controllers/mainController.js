angular.module('NewsScraperApp').controller('MainController',
    function ($scope, $interval, ApiService) {

        $scope.itemSelected = false;

        ApiService.getArticles().then(function (articles) {
            $scope.collectedArticles = articles.data.result;

            if (!$scope.itemSelected) {
                var article = $scope.collectedArticles[0];
                insertInnerPage(article.Link, article.ContentSelector, article.Charset);
            }
        });

        $scope.hasImage = function (article) {
            return article.Img != '';
        };

        $scope.openLink = function (link, contentSelector, charset) {
            $scope.itemSelected = true;
            $scope.$parent.selected = link;
            insertInnerPage(link, contentSelector, charset);
        };

        var insertInnerPage = function (link, contentSelector, charset) {
            $("#external-article").html('<object data="/read-html?url=' + link + '&selector=' + contentSelector + '&charset=' + charset + '"/>');
        };

        var updateInternalPageContent = $interval(function () {
            if (!_.isEmpty($('#external-article object'))) {
                $interval.cancel(updateInternalPageContent);
                $('#external-article object').css('height', calcHeight() + 'px');
            }
        }, 50);

        function calcHeight() {
            var originalHeight = "innerHeight" in window ? window.innerHeight : document.documentElement.offsetHeight;
            return originalHeight - 20;
        }

        function updateHeights(selectors) {
            _.each(selectors, function (selector) {
                $(selector).css('height', calcHeight() + 'px');
            });
        }

        function updateContentHeight() {
            updateHeights(['#left-column', '#right-column', '#external-article object']);
        }

        $(window).resize(function () {
            updateContentHeight();
        }).trigger("resize");
    });