// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.10
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.PolyaxonSdk);
  }
}(this, function(expect, PolyaxonSdk) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new PolyaxonSdk.TeamsV1Api();
  });

  describe('(package)', function() {
    describe('TeamsV1Api', function() {
      describe('createTeam', function() {
        it('should call createTeam successfully', function(done) {
          // TODO: uncomment, update parameter values for createTeam call and complete the assertions
          /*
          var owner = "owner_example";
          var body = new PolyaxonSdk.V1Team();
          body.uuid = "";
          body.name = "";
          body.projects = [""];
          body.created_at = new Date();
          body.updated_at = new Date();

          instance.createTeam(owner, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1Team);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            {
              let dataCtr = data.projects;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('createTeamMember', function() {
        it('should call createTeamMember successfully', function(done) {
          // TODO: uncomment, update parameter values for createTeamMember call and complete the assertions
          /*
          var owner = "owner_example";
          var team = "team_example";
          var body = new PolyaxonSdk.V1TeamMember();
          body.user = "";
          body.role = "";
          body.org_role = "";
          body.created_at = new Date();
          body.updated_at = new Date();

          instance.createTeamMember(owner, team, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1TeamMember);
            expect(data.user).to.be.a('string');
            expect(data.user).to.be("");
            expect(data.role).to.be.a('string');
            expect(data.role).to.be("");
            expect(data.org_role).to.be.a('string');
            expect(data.org_role).to.be("");
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('deleteTeam', function() {
        it('should call deleteTeam successfully', function(done) {
          // TODO: uncomment, update parameter values for deleteTeam call
          /*
          var owner = "owner_example";
          var team = "team_example";

          instance.deleteTeam(owner, team, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('deleteTeamMember', function() {
        it('should call deleteTeamMember successfully', function(done) {
          // TODO: uncomment, update parameter values for deleteTeamMember call
          /*
          var owner = "owner_example";
          var team = "team_example";
          var member_user = "member_user_example";
          var opts = {};
          opts.member_role = "member_role_example";
          opts.member_org_role = "member_org_role_example";
          opts.member_created_at = new Date("2013-10-20T19:20:30+01:00");
          opts.member_updated_at = new Date("2013-10-20T19:20:30+01:00");

          instance.deleteTeamMember(owner, team, member_user, opts, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('getTeam', function() {
        it('should call getTeam successfully', function(done) {
          // TODO: uncomment, update parameter values for getTeam call and complete the assertions
          /*
          var owner = "owner_example";
          var team = "team_example";

          instance.getTeam(owner, team, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1Team);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            {
              let dataCtr = data.projects;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('getTeamMember', function() {
        it('should call getTeamMember successfully', function(done) {
          // TODO: uncomment, update parameter values for getTeamMember call and complete the assertions
          /*
          var owner = "owner_example";
          var team = "team_example";
          var user = "user_example";

          instance.getTeamMember(owner, team, user, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1TeamMember);
            expect(data.user).to.be.a('string');
            expect(data.user).to.be("");
            expect(data.role).to.be.a('string');
            expect(data.role).to.be("");
            expect(data.org_role).to.be.a('string');
            expect(data.org_role).to.be("");
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('listTeamMembers', function() {
        it('should call listTeamMembers successfully', function(done) {
          // TODO: uncomment, update parameter values for listTeamMembers call and complete the assertions
          /*
          var owner = "owner_example";
          var team = "team_example";

          instance.listTeamMembers(owner, team, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ListTeamMembersResponse);
            expect(data.count).to.be.a('number');
            expect(data.count).to.be(0);
            {
              let dataCtr = data.results;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a(PolyaxonSdk.V1TeamMember);
                expect(data.user).to.be.a('string');
                expect(data.user).to.be("");
                expect(data.role).to.be.a('string');
                expect(data.role).to.be("");
                expect(data.org_role).to.be.a('string');
                expect(data.org_role).to.be("");
                expect(data.created_at).to.be.a(Date);
                expect(data.created_at).to.be(new Date());
                expect(data.updated_at).to.be.a(Date);
                expect(data.updated_at).to.be(new Date());
              }
            }
            expect(data.previous).to.be.a('string');
            expect(data.previous).to.be("");
            expect(data.next).to.be.a('string');
            expect(data.next).to.be("");

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('listTeamNames', function() {
        it('should call listTeamNames successfully', function(done) {
          // TODO: uncomment, update parameter values for listTeamNames call and complete the assertions
          /*
          var owner = "owner_example";
          var opts = {};
          opts.offset = 56;
          opts.limit = 56;
          opts.sort = "sort_example";
          opts.query = "query_example";

          instance.listTeamNames(owner, opts, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ListTeamsResponse);
            expect(data.count).to.be.a('number');
            expect(data.count).to.be(0);
            {
              let dataCtr = data.results;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a(PolyaxonSdk.V1Team);
                expect(data.uuid).to.be.a('string');
                expect(data.uuid).to.be("");
                expect(data.name).to.be.a('string');
                expect(data.name).to.be("");
                {
                  let dataCtr = data.projects;
                  expect(dataCtr).to.be.an(Array);
                  expect(dataCtr).to.not.be.empty();
                  for (let p in dataCtr) {
                    let data = dataCtr[p];
                    expect(data).to.be.a('string');
                    expect(data).to.be("");
                  }
                }
                expect(data.created_at).to.be.a(Date);
                expect(data.created_at).to.be(new Date());
                expect(data.updated_at).to.be.a(Date);
                expect(data.updated_at).to.be(new Date());
              }
            }
            expect(data.previous).to.be.a('string');
            expect(data.previous).to.be("");
            expect(data.next).to.be.a('string');
            expect(data.next).to.be("");

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('listTeams', function() {
        it('should call listTeams successfully', function(done) {
          // TODO: uncomment, update parameter values for listTeams call and complete the assertions
          /*
          var owner = "owner_example";
          var opts = {};
          opts.offset = 56;
          opts.limit = 56;
          opts.sort = "sort_example";
          opts.query = "query_example";

          instance.listTeams(owner, opts, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ListTeamsResponse);
            expect(data.count).to.be.a('number');
            expect(data.count).to.be(0);
            {
              let dataCtr = data.results;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a(PolyaxonSdk.V1Team);
                expect(data.uuid).to.be.a('string');
                expect(data.uuid).to.be("");
                expect(data.name).to.be.a('string');
                expect(data.name).to.be("");
                {
                  let dataCtr = data.projects;
                  expect(dataCtr).to.be.an(Array);
                  expect(dataCtr).to.not.be.empty();
                  for (let p in dataCtr) {
                    let data = dataCtr[p];
                    expect(data).to.be.a('string');
                    expect(data).to.be("");
                  }
                }
                expect(data.created_at).to.be.a(Date);
                expect(data.created_at).to.be(new Date());
                expect(data.updated_at).to.be.a(Date);
                expect(data.updated_at).to.be(new Date());
              }
            }
            expect(data.previous).to.be.a('string');
            expect(data.previous).to.be("");
            expect(data.next).to.be.a('string');
            expect(data.next).to.be("");

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('patchTeam', function() {
        it('should call patchTeam successfully', function(done) {
          // TODO: uncomment, update parameter values for patchTeam call and complete the assertions
          /*
          var owner = "owner_example";
          var team_name = "team_name_example";
          var body = new PolyaxonSdk.V1Team();
          body.uuid = "";
          body.name = "";
          body.projects = [""];
          body.created_at = new Date();
          body.updated_at = new Date();

          instance.patchTeam(owner, team_name, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1Team);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            {
              let dataCtr = data.projects;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('patchTeamMember', function() {
        it('should call patchTeamMember successfully', function(done) {
          // TODO: uncomment, update parameter values for patchTeamMember call and complete the assertions
          /*
          var owner = "owner_example";
          var team = "team_example";
          var member_user = "member_user_example";
          var body = new PolyaxonSdk.V1TeamMember();
          body.user = "";
          body.role = "";
          body.org_role = "";
          body.created_at = new Date();
          body.updated_at = new Date();

          instance.patchTeamMember(owner, team, member_user, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1TeamMember);
            expect(data.user).to.be.a('string');
            expect(data.user).to.be("");
            expect(data.role).to.be.a('string');
            expect(data.role).to.be("");
            expect(data.org_role).to.be.a('string');
            expect(data.org_role).to.be("");
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('updateTeam', function() {
        it('should call updateTeam successfully', function(done) {
          // TODO: uncomment, update parameter values for updateTeam call and complete the assertions
          /*
          var owner = "owner_example";
          var team_name = "team_name_example";
          var body = new PolyaxonSdk.V1Team();
          body.uuid = "";
          body.name = "";
          body.projects = [""];
          body.created_at = new Date();
          body.updated_at = new Date();

          instance.updateTeam(owner, team_name, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1Team);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            {
              let dataCtr = data.projects;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('updateTeamMember', function() {
        it('should call updateTeamMember successfully', function(done) {
          // TODO: uncomment, update parameter values for updateTeamMember call and complete the assertions
          /*
          var owner = "owner_example";
          var team = "team_example";
          var member_user = "member_user_example";
          var body = new PolyaxonSdk.V1TeamMember();
          body.user = "";
          body.role = "";
          body.org_role = "";
          body.created_at = new Date();
          body.updated_at = new Date();

          instance.updateTeamMember(owner, team, member_user, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1TeamMember);
            expect(data.user).to.be.a('string');
            expect(data.user).to.be("");
            expect(data.role).to.be.a('string');
            expect(data.role).to.be("");
            expect(data.org_role).to.be.a('string');
            expect(data.org_role).to.be("");
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
    });
  });

}));
