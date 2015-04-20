var warningTrackApp = angular.module('warningTrackApp', []);

warningTrackApp
  .filter('svgIconBaseHref', function ($sce) {
    return function(basesId) {
      var brId = 0;
      if (isNaN(Number(basesId))) {
        brId = basesId;
      }
      return $sce.trustAsResourceUrl('baseball-bases.svg#br' + brId);
    };
  });

warningTrackApp
  .filter('svgIconStrikesOutsHref', function ($sce) {
    return function(soId) {
      return $sce.trustAsResourceUrl('strikes-outs.svg#so' + soId);
    };
  });

warningTrackApp
  .filter('svgIconBallsHref', function ($sce) {
    return function(ballsId) {
      return $sce.trustAsResourceUrl('balls.svg#b' + ballsId);
    };
  });

warningTrackApp
  .filter('svgIconTeamsHref', function ($sce) {
    return function(teamAbbr) {
      return $sce.trustAsResourceUrl('teams.svg#' + teamAbbr);
    };
  });

warningTrackApp
  .filter('leverageToSvg', function ($sce) {
    return function(leverageIndex) {
      if (leverageIndex < 1.0) { warningId = "0"} else
      if (leverageIndex < 1.5) { warningId = "1"} else
      if (leverageIndex < 2.0) { warningId = "2"} else
      if (leverageIndex < 2.5) { warningId = "3"} else
      { warningId = "4"}
      return $sce.trustAsResourceUrl('warning.svg#w' + warningId);
    };
  });

warningTrackApp
  .filter('inningSuffix', function($filter) {
    var suffixes = ["th", "st", "nd", "rd"];
    return function(inning) {
      var relevantDigits = Number(inning) % 10;
      var suffix = (relevantDigits <= 3) ? suffixes[relevantDigits] : suffixes[0];
      return inning+suffix;
    };
  });


warningTrackApp.controller('WarningTrackCtrl', ['$scope', '$http', '$filter',
  function ($scope, $http, $filter) {
    $http.get('/games').success(function(data) {
      $scope.games = data;
    });
    $scope.orderProp = ['status.status','-leverage_index'];
  }]);
