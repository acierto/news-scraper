angular.module('NewsScraperApp').controller('MainController',
    function ($scope, $interval, ApiService) {

        $scope.itemSelected = false;

        ApiService.getArticles().then(function (articles) {
            $scope.newsSources = articles.data;

            if (!$scope.itemSelected) {
                $("#external-article").html('<object data="' + $scope.newsSources[0].Articles[0].Link + '"/>');
            }
        });

        $scope.hasImage = function (article) {
            return article.Img != '';
        };

        $scope.openLink = function (link) {
            $scope.itemSelected = true;
            $scope.$parent.selected = link;
            $("#external-article").html('<object data="' + link + '"/>');
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
            _.each(selectors, function(selector){
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