# Copyright 2017 Google Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from ftl.common import builder

import hashlib
import os
import subprocess
import tempfile

from containerregistry.client.v2_2 import append

_NODE_NAMESPACE = 'node-lock-cache'


class NodeApp(builder.JustApp):

    def CreatePackageBase(self, base_image, cache):
        if self._ctx.Contains('package-lock.json'):
            pl = self._ctx.GetFile('package-lock.json')
            checksum = hashlib.sha256(pl).hexdigest()
            print('Using package-lock.json for cache.')
        elif self._ctx.Contains('package-lock.json'):
            pj = self._ctx.GetFile('package.json')
            checksum = hashlib.sha256(pj).hexdigest()
            print('Using package.json for cache.')
        else:
            print('No package manifest found.')
            return base_image

        hit = cache.Get(base_image, _NODE_NAMESPACE, checksum)
        if hit:
            return hit

        tmp_dir = tempfile.mkdtemp()
        app_dir = os.path.join(tmp_dir, 'app')
        os.makedirs(app_dir)
        for p in ['package.json', 'package-lock.json']:
            if self._ctx.Contains(p):
                with open(os.path.join(app_dir, p), 'w') as f:
                    f.write(self._ctx.GetFile(p))

        print('Running npm install to generate package base.')
        subprocess.check_output(['npm', 'install', '.'], cwd=app_dir)
        subprocess.check_output([
            'tar', 'czf',
            '/workspace/node_modules.tar.gz', '.'], cwd=tmp_dir)

        with open('/workspace/node_modules.tar.gz', 'rb') as l:
            layer = l.read()

        with append.Layer(base_image, layer) as l:
            cache.Store(base_image, _NODE_NAMESPACE, checksum, l)
            return l
