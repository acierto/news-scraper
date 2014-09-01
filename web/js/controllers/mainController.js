angular.module('NewsScraperApp').controller('MainController',
    function ($scope, $interval, ApiService) {

        ApiService.getArticles().then(function (articles) {
            $scope.newsSources = articles.data;
        });

        $scope.hasImage = function (article) {
            return article.Img != '';
        };

        $scope.openLink = function (link) {
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

        function updateContentHeight() {
            $('#left-column').css('height', calcHeight() + 'px');
            $('#right-column').css('height', calcHeight() + 'px');
            $('#external-article object').css('height', calcHeight() + 'px');
        }

        $scope.$on('$routeChangeSuccess', function () {
            updateContentHeight();
        });

        $(window).resize(function () {
            updateContentHeight();
        }).trigger("resize");
    });