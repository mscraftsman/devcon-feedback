module.exports = {
  pwa: {
    workboxPluginMode: 'GenerateSW',
    workboxOptions: {
      swDest: 'service-worker.js',
      skipWaiting: true,
      clientsClaim: true,
    }
  }
}