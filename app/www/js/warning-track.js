var warningTrackApp = angular.module('warningTrackApp', ['ngCookies']);

warningTrackApp
  .filter('svgIconBaseHref', function($sce) {
    return function(basesId) {
      return $sce.trustAsResourceUrl('../img/baseball-bases.svg#br' + basesId);
    };
  });

warningTrackApp
  .filter('displayGameStatus', function() {
    return function(game) {
      var displayString = "";

      if (game.status.state < 3) {
        displayString = "Final";
      } else if (game.status.state == 3) {
        displayString = "Postponed";
      } else if (game.status.state > 10 && game.status.state < 20) {
        var d = new Date(game.date_time);

        var hr = d.getHours();
        var ampm = hr < 12 ? "am" : "pm";
        if (hr > 12) {
          hr = hr-12;
        }
        var min = d.getMinutes();
        if (min < 10) {
            min = "0" + min;
        }
        displayString = hr + ":" + min + " " + ampm;
      } else if (game.status.state == 21) {
        displayString = "Delayed";
      } else {
        var halfInning = "B ";
        if (game.status.half_inning == "Top") {
          halfInning = "T ";
        }
        displayString = halfInning + game.status.inning;
      }
      return displayString;
    };
  });

  warningTrackApp
    .filter('miniGameStatus', function() {
      return function(game) {
        var displayString = "";

        if (game.status.state < 3) {
          displayString = "Final";
        } else if (game.status.state == 3) {
          displayString = "PP";
        } else if (game.status.state > 10 && game.status.state < 20) {
          var d = new Date(game.date_time);

          var hr = d.getHours();
          var ampm = hr < 12 ? "am" : "pm";
          if (hr > 12) {
            hr = hr-12;
          }
          var min = d.getMinutes();
          if (min < 10) {
              min = "0" + min;
          }
          displayString = hr + ":" + min;
        } else if (game.status.state == 21) {
          displayString = "Delay";
        } else {
          var halfInning = "B";
          if (game.status.half_inning == "Top") {
            halfInning = "T";
          }
          displayString = halfInning + game.status.inning;
        }
        return displayString;
      };
    });

warningTrackApp
  .filter('svgIconStrikesOutsHref', function($sce) {
    return function(soId) {
      return $sce.trustAsResourceUrl('../img/strikes-outs.svg#so' + soId);
    };
  });

warningTrackApp
  .filter('svgIconBallsHref', function($sce) {
    return function(ballsId) {
      return $sce.trustAsResourceUrl('../img/balls.svg#b' + ballsId);
    };
  });

warningTrackApp
  .filter('svgIconTeamsHref', function($sce) {
    return function(teamAbbr) {
      return $sce.trustAsResourceUrl('../img/teams.svg#' + teamAbbr);
    };
  });

warningTrackApp
  .filter('leverageToSvg', function($sce) {
    return function(leverageIndex) {
      if (leverageIndex < 1.0) {
        warningId = "0"
      } else
      if (leverageIndex < 1.5) {
        warningId = "1"
      } else
      if (leverageIndex < 2.0) {
        warningId = "2"
      } else
      if (leverageIndex < 2.5) {
        warningId = "3"
      } else
      if (leverageIndex < 3.0) {
        warningId = "4"
      } else {
        warningId = "5"
      }
      return $sce.trustAsResourceUrl('../img/warning.svg#w' + warningId);
    };
  });

warningTrackApp
  .filter('logoPosition', function($filter) {
    return function(id) {
      return "{'background-image: url('../img/team_logos/" + yPos.toString() + ".svg'}";
    };
  });

warningTrackApp
  .filter('minilogoPosition', function($filter) {
    return function(id) {
      return "{'background-image: url('../img/team_logos/" + yPos.toString() + ".svg'}";
    };
  });

warningTrackApp
  .filter('gameDate', function(){
    return function(games) {
      var result = [];
      var today = new Date();
      var today_utc = new Date(today.getUTCFullYear(), today.getUTCMonth(), today.getUTCDate(), today.getUTCHours());
      var view_date = new Date(today.getUTCFullYear(), today.getUTCMonth(), today.getUTCDate());
      if (today_utc.getHours() < 12) {
        view_date = new Date(view_date.getTime() - 1000*60*60*24);
      }
      angular.forEach(games, function(game) {
        var game_date = new Date(game.date_time);
        if(game_date.getDate() == view_date.getDate() &&
           game_date.getMonth() == view_date.getMonth() &&
           game_date.getYear() == view_date.getYear()) {
          result.push(game);
        }
      });
      return result;
    };
  });

warningTrackApp
  .filter('sanitizeLink', function($sce){
    return function(link) {
      return $sce.trustAsResourceUrl(link);
    };
  });

warningTrackApp.controller('WarningTrackCtrl', ['$scope', '$http', '$filter', '$interval', '$cookieStore',
  function($scope, $http, $filter, $interval, $cookieStore) {
    $scope.orderProp = ['-status.leverage_index', '-status.state', 'date_time'];

    $scope.setCurrentGame = function(currGameId, currGameLink) {
      $cookieStore.put('currentGameId', currGameId);
      document.getElementById('mlbtv_iframe').src = currGameLink;
    }

    $scope.setTvModeType = function() {
      $cookieStore.put('tvModeType', $scope.items.name);
    }

    $scope.getTvModeType = function(type) {
      return $cookieStore.get('tvModeType');
    }

    $scope.changeGame = function() {
      var currentGameIdInt = 0;
      var tvModeType = "";
      currentGameIdInt = parseInt($cookieStore.get('currentGameId'));
      tvModeType = $cookieStore.get('tvModeType');

      var maxLi = 0;
      var maxLiGameLink = "http://mlb.tv";
      var maxGameId = 0;

      angular.forEach($scope.games, function(game) {
        if (game.status.leverage_index >= maxLi) {
          maxLiGameLink = game.links.mlb_tv;
          maxLi = game.status.leverage_index;
          maxGameId = game.id;
        }
      });

      var gameLink = maxLiGameLink;
      var gameId = maxGameId;

      angular.forEach($scope.games, function(game) {
        if (game.id == currentGameIdInt) {
          if ((maxLi < (game.status.leverage_index + 2) && (tvModeType == 'semi')) ||
             (tvModeType == 'none')) {
               gameLink = game.links.mlb_tv;
               gameId = game.id;
          }
        }
      });

      if (document.getElementById('mlbtv_iframe')) {
        var currentGameLink = document.getElementById('mlbtv_iframe').src;
        if (currentGameLink != gameLink) {
          $scope.setCurrentGame(gameId, gameLink);
        }
      }
    }

    $scope.modes = [{name: 'auto', value: 'auto'},
                    {name: 'semi', value: 'semi'},
                    {name: 'none', value: 'none'}];

    $http.get('/games').success(function(data) {
      $scope.games = data;
      $scope.changeGame();
    });

    $interval(function() {
      $http.get('/games').success(function(data) {
        if (document.getElementById('miniGamesContainer')) {
          document.getElementById('miniGamesContainer').scrollLeft = 0;
        }
        $scope.games = data;
        $scope.changeGame();
      });
    }, 30000);
  }
]);
