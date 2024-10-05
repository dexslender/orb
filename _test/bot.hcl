bot {
  token = env("BOT_TOKEN")
  guild = 0
  setup_commands = false
  global_commands = false
  log_level = info
}

activity_manager {
  enabled = true
  online_movil = true
  interval = duration("10s")
  activity "Hello!" { type = custom }
  activity "Here we go again" { type = custom }
}
