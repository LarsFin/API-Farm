using System.Collections.Generic;
using ApiFarm.Models.Impl;
using ApiFarm.Services;
using ApiFarm.Utils.Impl;
using Microsoft.AspNetCore.Mvc;

namespace ApiFarm.Controllers
{
    /// <summary>
    /// Controller for consuming requests to get, add, update or delete video game entities from storage.
    /// </summary>
    [NotSupportedExceptionFilter]
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
        /// Endpoint to retrieve all <see cref="VideoGame"/> models from storage.
        /// </summary>
        /// <returns>All video games in an array as JSON.</returns>
        [HttpGet]
        public ActionResult<IEnumerable<VideoGame>> GetAll()
        {
            var query = this.videoGameService.GetAll();
            return this.Ok(query.Result);
        }

        /// <summary>
        /// Endpoint to add a <see cref="VideoGame"/> to storage.
        /// </summary>
        /// <param name="videoGame">The <see cref="VideoGame"/> extracted from request body to be added to storage.</param>
        /// <returns>The video game which was added.</returns>
        [HttpPost]
        public ActionResult Add(VideoGame videoGame)
        {
            var query = this.videoGameService.Add(videoGame);

            if (query.Code != 0)
            {
                return this.BadRequest(query.Message);
            }

            videoGame = query.Result;
            return this.Created($"video_games/{videoGame.Id}", videoGame);
        }
    }
}
