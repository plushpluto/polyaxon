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
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/V1Replica'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./V1Replica'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.V1PytorchJob = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.V1Replica);
  }
}(this, function(ApiClient, V1Replica) {
  'use strict';

  /**
   * The V1PytorchJob model module.
   * @module model/V1PytorchJob
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>V1PytorchJob</code>.
   * @alias module:model/V1PytorchJob
   * @class
   */
  var exports = function() {
  };

  /**
   * Constructs a <code>V1PytorchJob</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/V1PytorchJob} obj Optional instance to populate.
   * @return {module:model/V1PytorchJob} The populated <code>V1PytorchJob</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('kind'))
        obj.kind = ApiClient.convertToType(data['kind'], 'String');
      if (data.hasOwnProperty('master'))
        obj.master = V1Replica.constructFromObject(data['master']);
      if (data.hasOwnProperty('worker'))
        obj.worker = V1Replica.constructFromObject(data['worker']);
    }
    return obj;
  }

  /**
   * @member {String} kind
   */
  exports.prototype.kind = undefined;

  /**
   * @member {module:model/V1Replica} master
   */
  exports.prototype.master = undefined;

  /**
   * @member {module:model/V1Replica} worker
   */
  exports.prototype.worker = undefined;

  return exports;

}));