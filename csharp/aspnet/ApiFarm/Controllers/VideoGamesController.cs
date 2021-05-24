using System.Collections.Generic;
using ApiFarm.Models.Impl;
using ApiFarm.Services;
using Microsoft.AspNetCore.Mvc;

namespace ApiFarm.Controllers
{
    /// <summary>
    /// Controller for consuming requests to get, add, update or delete video game entities from storage.
    /// </summary>
    [Route("video_games")]
    [ApiController]
    public class VideoGamesController : ControllerBase
    {
        private IService<VideoGame> videoGameService;

        /// <summary>
        /// Initializes a new instance of the <see cref="VideoGamesController"/> class.
        /// </summary>
        /// <param name="videoGameService">Service to query for creating, reading, updating and deleting <see cref="VideoGame"/> resources.</param>
        public VideoGamesController(IService<VideoGame> videoGameService)
        {
            this.videoGameService = videoGameService;
        }

        /// <summary>
        /// Endpoint to retrieve all video games from storage.
        /// </summary>
        /// <returns>All video games in an array as JSON.</returns>
        [HttpGet]
        public ActionResult<IEnumerable<VideoGame>> GetAll()
        {
            var query = this.videoGameService.GetAll();
            return this.Ok(query.Result);
        }
    }
}
