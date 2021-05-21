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
        /// <summary>
        /// Endpoint to retrieve all video games from storage.
        /// </summary>
        /// <returns>All video games in an array as JSON.</returns>
        [HttpGet]
        public string GetAll()
        {
            return "Hello World!";
        }
    }
}
