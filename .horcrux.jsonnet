local repo(name) = {
  name: name,
  source: 'https://git.jolheiser.com/' + name + '.git',
  dest: [
    'git@github.com:jolheiser/' + name,
    'git@tangled.sh:jolheiser.com/' + name,
  ],
};
{
  key: '~/.ssh/horcrux',
  interval: '15m',
  storage: '.horcrux',
  repos: [
    repo('horcrux'),
    repo('ugit'),
    repo('helix.drv'),
  ],
}
