local hc = import 'horcrux.libsonnet';
local repo(name) = {
  source: 'https://git.jolheiser.com/' + name + '.git',
  dest: [
    {
      forge: hc.GitHub('jolheiser', 'secrets/gh', name),
      url: 'https://github.com/jolheiser/' + name + '.git',
    },
    //{
    //  forge: hc.Gitea('jolheiser', 'secrets/gt', name),
    //  url: 'https://gitea.com/jolheiser/' + name + '.git',
    //},
  ],
};
{
  interval: '1h',
  storage: '.horcrux',
  repos: [
    repo('horcrux'),
  ],
}
