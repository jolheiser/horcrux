// Optionally import the jsonnet lib
local hc = import 'horcrux.libsonnet';

// Optional example of using jsonnet to remove some boilerplate
local repo(name) = {
  source: 'https://git.jolheiser.com/' + name + '.git',
  dest: [
    {
      forge: hc.GitHub('jolheiser', 'secret', name),
      url: 'https://github.com/jolheiser/' + name + '.git',
    },
    {
      forge: hc.Gitea('jolheiser', 'moreSecret', name),
      url: 'https://gitea.com/jolheiser/' + name + '.git',
    },
  ],
};

// Actual output config
{
  // https://pkg.go.dev/time#ParseDuration
  interval: '1h',
  storage: '.horcrux',
  repos: [
    repo('horcrux'),
    repo('ugit'),
    repo('helix.drv'),
  ],
}
