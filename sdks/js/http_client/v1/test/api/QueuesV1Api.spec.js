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
    instance = new PolyaxonSdk.QueuesV1Api();
  });

  describe('(package)', function() {
    describe('QueuesV1Api', function() {
      describe('createQueue', function() {
        it('should call createQueue successfully', function(done) {
          // TODO: uncomment, update parameter values for createQueue call and complete the assertions
          /*
          var owner = "owner_example";
          var agent = "agent_example";
          var body = new PolyaxonSdk.V1Queue();
          body.uuid = "";
          body.agent = "";
          body.name = "";
          body.priority = 0;
          body.concurrency = 0;
          body.created_at = new Date();
          body.updated_at = new Date();

          instance.createQueue(owner, agent, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1Agent);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.description).to.be.a('string');
            expect(data.description).to.be("");
            {
              let dataCtr = data.tags;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.disabled).to.be.a('boolean');
            expect(data.disabled).to.be(false);
            expect(data.deleted).to.be.a('boolean');
            expect(data.deleted).to.be(false);
            expect(data.namespace).to.be.a('string');
            expect(data.namespace).to.be("");
            expect(data.version_api).to.be.a(Object);
            expect(data.version_api).to.be();
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
      describe('deleteQueue', function() {
        it('should call deleteQueue successfully', function(done) {
          // TODO: uncomment, update parameter values for deleteQueue call
          /*
          var owner = "owner_example";
          var agent = "agent_example";
          var uuid = "uuid_example";

          instance.deleteQueue(owner, agent, uuid, function(error, data, response) {
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
      describe('getQueue', function() {
        it('should call getQueue successfully', function(done) {
          // TODO: uncomment, update parameter values for getQueue call and complete the assertions
          /*
          var owner = "owner_example";
          var agent = "agent_example";
          var uuid = "uuid_example";

          instance.getQueue(owner, agent, uuid, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1Queue);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.agent).to.be.a('string');
            expect(data.agent).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.priority).to.be.a('number');
            expect(data.priority).to.be(0);
            expect(data.concurrency).to.be.a('number');
            expect(data.concurrency).to.be(0);
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
      describe('listOrganizationQueueNames', function() {
        it('should call listOrganizationQueueNames successfully', function(done) {
          // TODO: uncomment, update parameter values for listOrganizationQueueNames call and complete the assertions
          /*
          var owner = "owner_example";
          var opts = {};
          opts.offset = 56;
          opts.limit = 56;
          opts.sort = "sort_example";
          opts.query = "query_example";

          instance.listOrganizationQueueNames(owner, opts, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ListQueuesResponse);
            expect(data.count).to.be.a('number');
            expect(data.count).to.be(0);
            {
              let dataCtr = data.results;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a(PolyaxonSdk.V1Queue);
                expect(data.uuid).to.be.a('string');
                expect(data.uuid).to.be("");
                expect(data.agent).to.be.a('string');
                expect(data.agent).to.be("");
                expect(data.name).to.be.a('string');
                expect(data.name).to.be("");
                expect(data.priority).to.be.a('number');
                expect(data.priority).to.be(0);
                expect(data.concurrency).to.be.a('number');
                expect(data.concurrency).to.be(0);
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
      describe('listOrganizationQueues', function() {
        it('should call listOrganizationQueues successfully', function(done) {
          // TODO: uncomment, update parameter values for listOrganizationQueues call and complete the assertions
          /*
          var owner = "owner_example";
          var opts = {};
          opts.offset = 56;
          opts.limit = 56;
          opts.sort = "sort_example";
          opts.query = "query_example";

          instance.listOrganizationQueues(owner, opts, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ListQueuesResponse);
            expect(data.count).to.be.a('number');
            expect(data.count).to.be(0);
            {
              let dataCtr = data.results;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a(PolyaxonSdk.V1Queue);
                expect(data.uuid).to.be.a('string');
                expect(data.uuid).to.be("");
                expect(data.agent).to.be.a('string');
                expect(data.agent).to.be("");
                expect(data.name).to.be.a('string');
                expect(data.name).to.be("");
                expect(data.priority).to.be.a('number');
                expect(data.priority).to.be(0);
                expect(data.concurrency).to.be.a('number');
                expect(data.concurrency).to.be(0);
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
      describe('listQueueNames', function() {
        it('should call listQueueNames successfully', function(done) {
          // TODO: uncomment, update parameter values for listQueueNames call and complete the assertions
          /*
          var owner = "owner_example";
          var agent = "agent_example";
          var opts = {};
          opts.offset = 56;
          opts.limit = 56;
          opts.sort = "sort_example";
          opts.query = "query_example";

          instance.listQueueNames(owner, agent, opts, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ListQueuesResponse);
            expect(data.count).to.be.a('number');
            expect(data.count).to.be(0);
            {
              let dataCtr = data.results;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a(PolyaxonSdk.V1Queue);
                expect(data.uuid).to.be.a('string');
                expect(data.uuid).to.be("");
                expect(data.agent).to.be.a('string');
                expect(data.agent).to.be("");
                expect(data.name).to.be.a('string');
                expect(data.name).to.be("");
                expect(data.priority).to.be.a('number');
                expect(data.priority).to.be(0);
                expect(data.concurrency).to.be.a('number');
                expect(data.concurrency).to.be(0);
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
      describe('listQueues', function() {
        it('should call listQueues successfully', function(done) {
          // TODO: uncomment, update parameter values for listQueues call and complete the assertions
          /*
          var owner = "owner_example";
          var agent = "agent_example";
          var opts = {};
          opts.offset = 56;
          opts.limit = 56;
          opts.sort = "sort_example";
          opts.query = "query_example";

          instance.listQueues(owner, agent, opts, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ListQueuesResponse);
            expect(data.count).to.be.a('number');
            expect(data.count).to.be(0);
            {
              let dataCtr = data.results;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a(PolyaxonSdk.V1Queue);
                expect(data.uuid).to.be.a('string');
                expect(data.uuid).to.be("");
                expect(data.agent).to.be.a('string');
                expect(data.agent).to.be("");
                expect(data.name).to.be.a('string');
                expect(data.name).to.be("");
                expect(data.priority).to.be.a('number');
                expect(data.priority).to.be(0);
                expect(data.concurrency).to.be.a('number');
                expect(data.concurrency).to.be(0);
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
      describe('patchQueue', function() {
        it('should call patchQueue successfully', function(done) {
          // TODO: uncomment, update parameter values for patchQueue call and complete the assertions
          /*
          var owner = "owner_example";
          var queue_agent = "queue_agent_example";
          var queue_uuid = "queue_uuid_example";
          var body = new PolyaxonSdk.V1Queue();
          body.uuid = "";
          body.agent = "";
          body.name = "";
          body.priority = 0;
          body.concurrency = 0;
          body.created_at = new Date();
          body.updated_at = new Date();

          instance.patchQueue(owner, queue_agent, queue_uuid, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1Queue);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.agent).to.be.a('string');
            expect(data.agent).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.priority).to.be.a('number');
            expect(data.priority).to.be(0);
            expect(data.concurrency).to.be.a('number');
            expect(data.concurrency).to.be(0);
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
      describe('updateQueue', function() {
        it('should call updateQueue successfully', function(done) {
          // TODO: uncomment, update parameter values for updateQueue call and complete the assertions
          /*
          var owner = "owner_example";
          var queue_agent = "queue_agent_example";
          var queue_uuid = "queue_uuid_example";
          var body = new PolyaxonSdk.V1Queue();
          body.uuid = "";
          body.agent = "";
          body.name = "";
          body.priority = 0;
          body.concurrency = 0;
          body.created_at = new Date();
          body.updated_at = new Date();

          instance.updateQueue(owner, queue_agent, queue_uuid, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1Queue);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.agent).to.be.a('string');
            expect(data.agent).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.priority).to.be.a('number');
            expect(data.priority).to.be(0);
            expect(data.concurrency).to.be.a('number');
            expect(data.concurrency).to.be(0);
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
