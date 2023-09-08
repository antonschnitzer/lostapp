export default {
  auth: {
    session: 'auth/session',
    login: 'auth/discord/login',
    logout: 'auth/discord/logout',
  },
  players: {
    all: 'players',
    byTag: 'players/:tag',
    leaderboard: 'players/leaderboard/:statsId',
    comparableStats: 'players/comparable-stats',
  },
  clans: {
    all: 'clans',
    byTag: 'clans/:tag',
    leading: 'clans/leading',
    settings: 'clans/:tag/settings',
    members: {
      kickpoints: {
        all: 'clans/:clanTag/members/kickpoints',
        byTag: 'clans/:clanTag/members/:memberTag/kickpoints',
      },
    },
  },
  notifications: 'notifications',
} as const
