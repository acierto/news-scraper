angular.module('NewsScraperApp').controller('MainController',
    function ($scope, $interval, ApiService) {

        $scope.itemSelected = false;

        ApiService.getArticles().then(function (articles) {
            $scope.groupedArticles = articles.data;

            if (!$scope.itemSelected) {
                var group = $scope.groupedArticles[0];
                insertInnerPage(group.Articles[0].Link, group.ContentSelector);
            }
        });

        $scope.hasImage = function (article) {
            return article.Img != '';
        };

        $scope.openLink = function (link, contentSelector) {
            $scope.itemSelected = true;
            $scope.$parent.selected = link;
            insertInnerPage(link, contentSelector);
        };

        var insertInnerPage = function (link, contentSelector) {
            $("#external-article").html('<object data="/read-html?url=' + link + '&selector=' + contentSelector + ' "/>');
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