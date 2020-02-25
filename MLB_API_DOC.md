# Stats API MLB
# Description
For a long time I used mlb's legacy api. But they have shut that down and moved to statsapi.mlb.com
I will try to document the endpoints I need for warning-track here.

## Host
statsapi.mlb.com

## Endpoints
### Schedule
#### URI
/api/v1/schedule
?language=en
&sportId=1
&date=02/24/2020
&sortBy=gameDate
&hydrate=game(content(summary,media(epg))),linescore(runners),flags,team

#### Example Response
```
{
  "copyright" : "Copyright 2020 MLB Advanced Media, L.P.  Use of any content on this page acknowledges agreement to the terms posted here http://gdx.mlb.com/components/copyright.txt",
  "totalItems" : 16,
  "totalEvents" : 0,
  "totalGames" : 16,
  "totalGamesInProgress" : 0,
  "dates" : [ {
    "date" : "2020-02-24",
    "totalItems" : 16,
    "totalEvents" : 0,
    "totalGames" : 16,
    "totalGamesInProgress" : 0,
    "games" : [ {
      "gamePk" : 605699,
      "link" : "/api/v1.1/game/605699/feed/live",
      "gameType" : "S",
      "season" : "2020",
      "gameDate" : "2020-02-24T18:05:00Z",
      "status" : {
        "abstractGameState" : "Final",
        "codedGameState" : "F",
        "detailedState" : "Final",
        "statusCode" : "F",
        "abstractGameCode" : "F"
      },
      "teams" : {
        "away" : {
          "leagueRecord" : {
            "wins" : 1,
            "losses" : 2,
            "pct" : ".333"
          },
          "score" : 7,
          "team" : {
            "id" : 110,
            "name" : "Baltimore Orioles",
            "link" : "/api/v1/teams/110",
            "season" : 2020,
            "venue" : {
              "id" : 2,
              "name" : "Oriole Park at Camden Yards",
              "link" : "/api/v1/venues/2"
            },
            "teamCode" : "bal",
            "fileCode" : "bal",
            "abbreviation" : "BAL",
            "teamName" : "Orioles",
            "locationName" : "Baltimore",
            "firstYearOfPlay" : "1901",
            "league" : {
              "id" : 103,
              "name" : "American League",
              "link" : "/api/v1/league/103"
            },
            "division" : {
              "id" : 201,
              "name" : "American League East",
              "link" : "/api/v1/divisions/201"
            },
            "sport" : {
              "id" : 1,
              "link" : "/api/v1/sports/1",
              "name" : "Major League Baseball"
            },
            "shortName" : "Baltimore",
            "springLeague" : {
              "id" : 115,
              "name" : "Grapefruit League",
              "link" : "/api/v1/league/115",
              "abbreviation" : "GL"
            },
            "allStarStatus" : "N",
            "active" : true
          },
          "splitSquad" : false,
          "seriesNumber" : 3,
          "springLeague" : {
            "id" : 115,
            "name" : "Grapefruit League",
            "link" : "/api/v1/league/115",
            "abbreviation" : "GL"
          }
        },
        "home" : {
          "leagueRecord" : {
            "wins" : 2,
            "losses" : 0,
            "pct" : "1.000"
          },
          "score" : 8,
          "team" : {
            "id" : 143,
            "name" : "Philadelphia Phillies",
            "link" : "/api/v1/teams/143",
            "season" : 2020,
            "venue" : {
              "id" : 2681,
              "name" : "Citizens Bank Park",
              "link" : "/api/v1/venues/2681"
            },
            "teamCode" : "phi",
            "fileCode" : "phi",
            "abbreviation" : "PHI",
            "teamName" : "Phillies",
            "locationName" : "Philadelphia",
            "firstYearOfPlay" : "1883",
            "league" : {
              "id" : 104,
              "name" : "National League",
              "link" : "/api/v1/league/104"
            },
            "division" : {
              "id" : 204,
              "name" : "National League East",
              "link" : "/api/v1/divisions/204"
            },
            "sport" : {
              "id" : 1,
              "link" : "/api/v1/sports/1",
              "name" : "Major League Baseball"
            },
            "shortName" : "Philadelphia",
            "springLeague" : {
              "id" : 115,
              "name" : "Grapefruit League",
              "link" : "/api/v1/league/115",
              "abbreviation" : "GL"
            },
            "allStarStatus" : "N",
            "active" : true
          },
          "splitSquad" : false,
          "seriesNumber" : 3,
          "springLeague" : {
            "id" : 115,
            "name" : "Grapefruit League",
            "link" : "/api/v1/league/115",
            "abbreviation" : "GL"
          }
        }
      },
      "linescore" : {
        "note" : "One out when winning run scored.",
        "currentInning" : 9,
        "currentInningOrdinal" : "9th",
        "inningState" : "Bottom",
        "inningHalf" : "Bottom",
        "isTopInning" : false,
        "scheduledInnings" : 9,
        "innings" : [ {
          "num" : 1,
          "ordinalNum" : "1st",
          "home" : {
            "runs" : 0,
            "hits" : 1,
            "errors" : 0,
            "leftOnBase" : 1
          },
          "away" : {
            "runs" : 0,
            "hits" : 0,
            "errors" : 0,
            "leftOnBase" : 1
          }
        }, {
          "num" : 2,
          "ordinalNum" : "2nd",
          "home" : {
            "runs" : 0,
            "hits" : 1,
            "errors" : 0,
            "leftOnBase" : 0
          },
          "away" : {
            "runs" : 0,
            "hits" : 1,
            "errors" : 0,
            "leftOnBase" : 0
          }
        }, {
          "num" : 3,
          "ordinalNum" : "3rd",
          "home" : {
            "runs" : 0,
            "hits" : 0,
            "errors" : 0,
            "leftOnBase" : 1
          },
          "away" : {
            "runs" : 0,
            "hits" : 0,
            "errors" : 0,
            "leftOnBase" : 0
          }
        }, {
          "num" : 4,
          "ordinalNum" : "4th",
          "home" : {
            "runs" : 3,
            "hits" : 4,
            "errors" : 0,
            "leftOnBase" : 1
          },
          "away" : {
            "runs" : 0,
            "hits" : 1,
            "errors" : 0,
            "leftOnBase" : 1
          }
        }, {
          "num" : 5,
          "ordinalNum" : "5th",
          "home" : {
            "runs" : 0,
            "hits" : 0,
            "errors" : 0,
            "leftOnBase" : 0
          },
          "away" : {
            "runs" : 0,
            "hits" : 0,
            "errors" : 0,
            "leftOnBase" : 0
          }
        }, {
          "num" : 6,
          "ordinalNum" : "6th",
          "home" : {
            "runs" : 4,
            "hits" : 5,
            "errors" : 1,
            "leftOnBase" : 2
          },
          "away" : {
            "runs" : 1,
            "hits" : 2,
            "errors" : 0,
            "leftOnBase" : 1
          }
        }, {
          "num" : 7,
          "ordinalNum" : "7th",
          "home" : {
            "runs" : 0,
            "hits" : 0,
            "errors" : 0,
            "leftOnBase" : 1
          },
          "away" : {
            "runs" : 2,
            "hits" : 2,
            "errors" : 0,
            "leftOnBase" : 0
          }
        }, {
          "num" : 8,
          "ordinalNum" : "8th",
          "home" : {
            "runs" : 0,
            "hits" : 0,
            "errors" : 1,
            "leftOnBase" : 0
          },
          "away" : {
            "runs" : 4,
            "hits" : 5,
            "errors" : 0,
            "leftOnBase" : 1
          }
        }, {
          "num" : 9,
          "ordinalNum" : "9th",
          "home" : {
            "runs" : 1,
            "hits" : 2,
            "errors" : 0,
            "leftOnBase" : 1
          },
          "away" : {
            "runs" : 0,
            "hits" : 0,
            "errors" : 1,
            "leftOnBase" : 0
          }
        } ],
        "teams" : {
          "home" : {
            "runs" : 8,
            "hits" : 13,
            "errors" : 2,
            "leftOnBase" : 7
          },
          "away" : {
            "runs" : 7,
            "hits" : 11,
            "errors" : 1,
            "leftOnBase" : 4
          }
        },
        "defense" : {
          "batter" : {
            "id" : 602922,
            "fullName" : "Jose Rondon",
            "link" : "/api/v1/people/602922"
          },
          "onDeck" : {
            "id" : 663630,
            "fullName" : "Ryan McKenna",
            "link" : "/api/v1/people/663630"
          },
          "inHole" : {
            "id" : 669200,
            "fullName" : "Mason McCoy",
            "link" : "/api/v1/people/669200"
          }
        },
        "offense" : {
          "first" : {
            "id" : 663897,
            "fullName" : "Luke Williams",
            "link" : "/api/v1/people/663897"
          }
        },
        "balls" : 0,
        "strikes" : 0,
        "outs" : 1
      },
      "venue" : {
        "id" : 2700,
        "name" : "Spectrum Field",
        "link" : "/api/v1/venues/2700"
      },
      "content" : {
        "link" : "/api/v1/game/605699/content",
        "editorial" : { },
        "media" : {
          "epg" : [ {
            "title" : "MLBTV",
            "items" : [ {
              "id" : 111429107,
              "contentId" : "05d62c94-bd91-48ad-9f0e-f9ef12ba75f9",
              "mediaId" : "9243724d-0ad8-434c-bc91-d1db9a92e2c3",
              "mediaState" : "MEDIA_ARCHIVE",
              "mediaFeedType" : "HOME",
              "mediaFeedSubType" : "143",
              "callLetters" : "NBCSP",
              "foxAuthRequired" : false,
              "tbsAuthRequired" : false,
              "espnAuthRequired" : false,
              "fs1AuthRequired" : false,
              "mlbnAuthRequired" : false,
              "freeGame" : false
            } ]
          }, {
            "title" : "MLBTV-Audio",
            "items" : [ {
              "id" : 111429107,
              "type" : "",
              "mediaFeedType" : "HOME",
              "description" : "NBCSP",
              "renditionName" : "English",
              "language" : "EN"
            } ]
          }, {
            "title" : "Audio",
            "items" : [ ]
          } ],
          "epgAlternate" : [ {
            "items" : [ ],
            "title" : "Extended Highlights"
          }, {
            "items" : [ {
              "type" : "video",
              "state" : "A",
              "date" : "2020-02-24T23:02:22.045Z",
              "id" : "orioles-vs-phillies-recap-2-24",
              "headline" : "Orioles vs. Phillies Recap 2/24",
              "seoTitle" : "",
              "slug" : "orioles-vs-phillies-recap-2-24",
              "blurb" : "Phillies crack four homers in walk off win | 2/24/20",
              "keywordsAll" : [ {
                "type" : "game",
                "value" : "gamepk-605699",
                "displayName" : "2020/02/24 bal@phi"
              }, {
                "type" : "game_pk",
                "value" : "605699",
                "displayName" : "2020/02/24 bal@phi"
              }, {
                "type" : "team",
                "value" : "teamid-143",
                "displayName" : "Philadelphia Phillies"
              }, {
                "type" : "team_id",
                "value" : "143",
                "displayName" : "Philadelphia Phillies"
              }, {
                "type" : "team",
                "value" : "teamid-110",
                "displayName" : "Baltimore Orioles"
              }, {
                "type" : "team_id",
                "value" : "110",
                "displayName" : "Baltimore Orioles"
              }, {
                "type" : "taxonomy",
                "value" : "send-to-news-mlb-feed",
                "displayName" : "Send To News MLB feed"
              }, {
                "type" : "taxonomy",
                "value" : "game-recap",
                "displayName" : "game recap"
              }, {
                "type" : "mlbtax",
                "value" : "mlb_recap",
                "displayName" : "Game Recap"
              }, {
                "type" : "subject",
                "value" : "MLBCOM_GAME_RECAP",
                "displayName" : "MLBCOM_GAME_RECAP"
              }, {
                "type" : "taxonomy",
                "value" : "spring-training",
                "displayName" : "Spring Training"
              }, {
                "type" : "taxonomy",
                "value" : "cactus-league",
                "displayName" : "Cactus League"
              }, {
                "type" : "taxonomy",
                "value" : "eclat-feed",
                "displayName" : "Eclat feed"
              }, {
                "type" : "taxonomy",
                "value" : "international-feed",
                "displayName" : "International Partner feed"
              } ],
              "keywordsDisplay" : [ {
                "type" : "mlbtax",
                "value" : "mlb_recap",
                "displayName" : "Game Recap"
              }, {
                "type" : "subject",
                "value" : "MLBCOM_GAME_RECAP",
                "displayName" : "MLBCOM_GAME_RECAP"
              } ],
              "image" : {
                "title" : "bal-vs-phi",
                "altText" : null,
                "cuts" : [ {
                  "aspectRatio" : "16:9",
                  "width" : 1920,
                  "height" : 1080,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_1920,h_1080,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_1920,h_1080,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_1920,h_1080,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 1440,
                  "height" : 810,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_1440,h_810,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_1440,h_810,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_1440,h_810,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 1280,
                  "height" : 720,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_1280,h_720,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_1280,h_720,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_1280,h_720,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 960,
                  "height" : 540,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_960,h_540,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_960,h_540,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_960,h_540,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 800,
                  "height" : 448,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_800,h_448,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_800,h_448,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_800,h_448,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 720,
                  "height" : 405,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_720,h_405,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_720,h_405,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_720,h_405,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 684,
                  "height" : 385,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_684,h_385,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_684,h_385,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_684,h_385,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 640,
                  "height" : 360,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_640,h_360,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_640,h_360,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_640,h_360,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 496,
                  "height" : 279,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_496,h_279,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_496,h_279,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_496,h_279,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 480,
                  "height" : 270,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_480,h_270,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_480,h_270,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_480,h_270,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 430,
                  "height" : 242,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_430,h_242,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_430,h_242,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_430,h_242,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 400,
                  "height" : 224,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_400,h_224,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_400,h_224,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_400,h_224,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 320,
                  "height" : 180,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_320,h_180,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_320,h_180,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_320,h_180,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 270,
                  "height" : 154,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_270,h_154,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_270,h_154,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_270,h_154,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 248,
                  "height" : 138,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_248,h_138,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_248,h_138,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_248,h_138,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 215,
                  "height" : 121,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_215,h_121,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_215,h_121,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_215,h_121,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 209,
                  "height" : 118,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_209,h_118,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_209,h_118,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_209,h_118,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 135,
                  "height" : 77,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_135,h_77,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_135,h_77,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_135,h_77,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "16:9",
                  "width" : 124,
                  "height" : 70,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_124,h_70,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_124,h_70,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_124,h_70,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "4:3",
                  "width" : 222,
                  "height" : 168,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_222,h_168,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_222,h_168,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_222,h_168,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "4:3",
                  "width" : 192,
                  "height" : 144,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_192,h_144,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_192,h_144,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_192,h_144,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "4:3",
                  "width" : 148,
                  "height" : 112,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_148,h_112,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_148,h_112,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_148,h_112,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "4:3",
                  "width" : 96,
                  "height" : 72,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_96,h_72,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_96,h_72,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_96,h_72,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "4:3",
                  "width" : 74,
                  "height" : 56,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_74,h_56,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_74,h_56,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_74,h_56,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                }, {
                  "aspectRatio" : "64:27",
                  "width" : 1920,
                  "height" : 810,
                  "src" : "https://img.mlbstatic.com/mlb-images/image/private/w_1920,h_810,f_jpg,c_fill,g_auto/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at2x" : "https://img.mlbstatic.com/mlb-images/image/private/w_1920,h_810,f_jpg,c_fill,g_auto,dpr_2.0/mlb/hyzg2quc2vrb6xaofqhk.jpg",
                  "at3x" : "https://img.mlbstatic.com/mlb-images/image/private/w_1920,h_810,f_jpg,c_fill,g_auto,dpr_3.0/mlb/hyzg2quc2vrb6xaofqhk.jpg"
                } ]
              },
              "noIndex" : false,
              "mediaPlaybackId" : "bc833f7e-561939f4-bb260abc-csvm-diamondx64-asset",
              "title" : "Orioles vs. Phillies Recap 2/24",
              "description" : "Mikie Mahtook's three-run home run was one of four the Phillies hit in their 8-7 walk-off win over the Orioles in Spring Training",
              "duration" : "00:02:31",
              "mediaPlaybackUrl" : "",
              "playbacks" : [ {
                "name" : "mp4Avc",
                "url" : "https://cuts.diamond.mlb.com/FORGE/2020/2020-02/24/bc833f7e-561939f4-bb260abc-csvm-diamondx64-asset_1280x720_59_4000K.mp4",
                "width" : "",
                "height" : ""
              }, {
                "name" : "hlsCloud",
                "url" : "https://cuts.diamond.mlb.com/FORGE/2020/2020-02/24/bc833f7e-561939f4-bb260abc-csvm-diamondx64-asset.m3u8",
                "width" : "",
                "height" : ""
              }, {
                "name" : "HTTP_CLOUD_WIRED",
                "url" : "https://cuts.diamond.mlb.com/FORGE/2020/2020-02/24/bc833f7e-561939f4-bb260abc-csvm-diamondx64-asset.m3u8",
                "width" : "",
                "height" : ""
              }, {
                "name" : "HTTP_CLOUD_WIRED_60",
                "url" : "https://cuts.diamond.mlb.com/FORGE/2020/2020-02/24/bc833f7e-561939f4-bb260abc-csvm-diamondx64-asset.m3u8",
                "width" : "",
                "height" : ""
              }, {
                "name" : "highBit",
                "url" : "https://cuts.diamond.mlb.com/FORGE/2020/2020-02/24/bc833f7e-561939f4-bb260abc-csvm-diamondx64-asset_1280x720_59_16000K.mp4",
                "width" : "",
                "height" : ""
              } ]
            } ],
            "title" : "Daily Recap"
          } ],
          "freeGame" : false,
          "enhancedGame" : false
        },
        "highlights" : { },
        "summary" : {
          "hasPreviewArticle" : false,
          "hasRecapArticle" : false,
          "hasWrapArticle" : false,
          "hasHighlightsVideo" : true
        },
        "gameNotes" : { }
      },
      "gameNumber" : 1,
      "publicFacing" : true,
      "doubleHeader" : "N",
      "gamedayType" : "Y",
      "tiebreaker" : "N",
      "calendarEventID" : "14-605699-2020-02-24",
      "seasonDisplay" : "2020",
      "dayNight" : "day",
      "scheduledInnings" : 9,
      "inningBreakLength" : 0,
      "gamesInSeries" : 1,
      "seriesGameNumber" : 1,
      "seriesDescription" : "Spring Training",
      "flags" : {
        "noHitter" : false,
        "perfectGame" : false
      },
      "recordSource" : "S",
      "ifNecessary" : "N",
      "ifNecessaryDescription" : "Normal Game"
    }, ... ],
    "events" : [ ]
  } ]
}
```
